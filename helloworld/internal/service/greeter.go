package service

import (
	"context"
	"fmt"
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
	"helloworld/internal/data"
	"helloworld/pkg"
	"log"
	"math/rand"
	"regexp"
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

	c.db.GetRedis().Del(context.Background(), "SendSms"+in.Mobile)

	token, _ := pkg.NewJWT("2211a").CreateToken(pkg.CustomClaims{
		ID: uint(user.Id),
	})

	return &v1.LoginReply{
		Code:    200,
		Message: "登录成功",
		Token:   token,
	}, nil
}

func (c *GreeterService) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {

	var user biz.User
	if err := c.db.GetDB().Where("id = ?", in.Id).Find(&user).Error; err != nil {
		return &v1.UpdateUserReply{
			Code:    500,
			Message: "用户不存在",
		}, nil
	}

	users := biz.User{
		Id:       int32(in.Id),
		Mobile:   in.Mobile,
		NickName: in.NickName,
		Sex:      in.Sex,
	}
	if !isValidMobile(in.Mobile) {
		return &v1.UpdateUserReply{
			Code:    500,
			Message: "手机号格式错误",
		}, nil
	}

	if !pkg.TextCensor(in.NickName) {
		return &v1.UpdateUserReply{
			Code:    500,
			Message: "昵称含有违规词",
		}, nil
	}

	if err := c.db.GetDB().Model(&user).Updates(&users).Error; err != nil {
		return &v1.UpdateUserReply{
			Code:    500,
			Message: "完善信息失败",
		}, nil
	}

	return &v1.UpdateUserReply{
		Code:    200,
		Message: "完善信息成功",
	}, nil
}

func (c *GreeterService) InfoUser(ctx context.Context, in *v1.InfoUserRequest) (*v1.InfoUserReply, error) {
	var user biz.User
	if err := c.db.GetDB().Where("id = ?", in.Id).Find(&user).Error; err != nil {
		return &v1.InfoUserReply{
			Code:    500,
			Message: "用户不存在",
		}, nil
	}

	var users biz.User

	if err := c.db.GetDB().Find(&users).Error; err != nil {
		return &v1.InfoUserReply{
			Code:    500,
			Message: "个人信息展示失败",
		}, nil
	}

	return &v1.InfoUserReply{
		Code:    200,
		Message: "个人信息展示成功",
		Info: []*v1.Data{
			{
				Mobile:   user.Mobile,
				UserName: user.UserName,
				NickName: user.NickName,
				Sex:      user.Sex,
			},
		},
	}, nil

}

func (c *GreeterService) RealName(ctx context.Context, in *v1.RealNameRequest) (*v1.RealNameReply, error) {

	if in.UserId == 0 || in.RealName == "" || in.IdCard == "" {
		return &v1.RealNameReply{
			Code:    500,
			Message: "参数不能为空",
		}, nil
	}

	var user biz.User
	err := c.db.GetDB().Where("id = ?", in.UserId).First(&user).Error
	if err != nil {
		return &v1.RealNameReply{
			Code:    500,
			Message: "用户不存在",
		}, nil
	}

	name := pkg.Realname(in.RealName, in.IdCard)

	if err := c.db.GetDB().Model(&user).Updates(name).Error; err != nil {
		return &v1.RealNameReply{
			Code:    500,
			Message: "实名认证失败",
		}, nil
	}

	return &v1.RealNameReply{
		Code:    200,
		Message: "实名认证成功",
	}, nil
}
