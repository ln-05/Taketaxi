package service

import (
	"context"
	"fmt"
	"helloworld/internal/data"
	"helloworld/pkg"

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
	// 临时放宽验证规则，只检查长度和是否都是数字
	if len(mobile) != 11 {
		fmt.Printf("手机号长度不正确: %d\n", len(mobile))
		return false
	}

	// 检查是否都是数字
	for _, char := range mobile {
		if char < '0' || char > '9' {
			fmt.Printf("手机号包含非数字字符: %c\n", char)
			return false
		}
	}

	// 检查是否以1开头
	if mobile[0] != '1' {
		fmt.Printf("手机号不是以1开头\n")
		return false
	}

	fmt.Printf("手机号验证通过: %s\n", mobile)
	return true
}

func (c *GreeterService) SendSms(_ context.Context, in *v1.SendSmsRequest) (*v1.SendSmsReply, error) {
	fmt.Printf("收到手机号: %s, 长度: %d\n", in.Mobile, len(in.Mobile))

	if !isValidMobile(in.Mobile) {
		return nil, fmt.Errorf("手机号格式不正确")
	}
	// 用redis分布式锁实现60秒内只能获取一次
	set, _ := c.db.GetRedis().SetNX(context.Background(), "SendSmsLimit:"+in.Mobile, 1, time.Second*60).Result()
	if !set {
		return &v1.SendSmsReply{
			Code:    200,
			Message: "短信发送过于频繁，请60秒后再试",
		}, nil
	}

	//生成验证码
	code := rand.Intn(9000) + 1000

	//将生成的验证码存储到redis缓存中，并设置过期时间5分钟
	c.db.GetRedis().Set(context.Background(), "SendSms"+in.Mobile, code, time.Minute*5)

	return &v1.SendSmsReply{
		Code:    200,
		Message: "发送验证码成功",
	}, nil
}
func (c *GreeterService) Login(_ context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	//判断手机号
	if !isValidMobile(in.Mobile) {
		return nil, fmt.Errorf("手机号格式不正确")
	}

	var user biz.User
	c.db.GetDB().Where("mobile=?", in.Mobile).Find(&user)
	// 判断用户是否存在，如果不存在直接注册

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
	// 接受验证码，判断验证码是否过期和验证码是否错误
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
