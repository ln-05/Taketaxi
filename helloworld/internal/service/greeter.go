package service

import (
	"context"
	"fmt"
	"helloworld/internal/data"
	"helloworld/pkg"
	"log"
	"regexp"

	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
	"math/rand"
	"time"
)

// GreeterService is a greeter service.
type GreeterService struct {
	db *data.Data
	v1.UnimplementedGreeterServer
	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, db *data.Data) *GreeterService {
	return &GreeterService{
		uc: uc,
		db: db,
	}
}

func isValidMobile(mobile string) bool {
	pattern := `^1[3-9]\d{9}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(mobile)
}

func (c *GreeterService) SendSms(_ context.Context, in *v1.SendSmsRequest) (*v1.SendSmsReply, error) {

	if !isValidMobile(in.Mobile) {
		return nil, fmt.Errorf("手机号格式不正确")
	}

	set, _ := c.db.GetRedis().SetNX(context.Background(), "SendSmsLimit:"+in.Mobile, 1, time.Second*60).Result()
	if !set {
		return &v1.SendSmsReply{
			Code:    200,
			Message: "短信发送过于频繁，请60秒后再试",
		}, nil
	}

	code := rand.Intn(900000) + 100000

	c.db.GetRedis().Set(context.Background(), "SendSms"+in.Mobile, code, time.Minute*5)

	log.Println(code)
	return &v1.SendSmsReply{
		Code:    200,
		Message: "发送验证码成功",
	}, nil
}

func (c *GreeterService) Login(_ context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {

	if !isValidMobile(in.Mobile) {
		return nil, fmt.Errorf("手机号格式不正确")
	}

	var user biz.User
	c.db.GetDB().Where("mobile=?", in.Mobile).Find(&user)

	if user.Id == 0 {
		newUser := biz.User{
			Mobile:   in.Mobile,
			UserName: "张三",
			NickName: "用户" + in.Mobile,
			Sex:      "默认",
		}

		if err := c.db.GetDB().Create(&newUser).Error; err != nil {
			return nil, fmt.Errorf("注册失败")
		}
	}

	get := c.db.GetRedis().Get(context.Background(), "SendSms"+in.Mobile)
	if get.Err() != nil {
		return nil, fmt.Errorf("验证码已过期")
	}

	if get.Val() != in.SendSmsCode {
		return nil, fmt.Errorf("验证码错误")
	}
	// 5.验证成功后，删除验证码
	c.db.GetRedis().Del(context.Background(), "SendSms"+in.Mobile)

	token, _ := pkg.NewJWT("2210a").CreateToken(pkg.CustomClaims{
		ID: uint(user.Id),
	})

	return &v1.LoginReply{
		Code:    200,
		Message: "登录成功",
		Token:   token,
	}, nil
}
