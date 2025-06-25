package service

import (
	"context"
	"fmt"
	"helloworld/internal/pkg"
	"math/rand"
	"regexp"
	"time"

	"gorm.io/gorm"

	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"

	"github.com/go-redis/redis/v8"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc  *biz.GreeterUsecase
	Db  *gorm.DB
	Rdb *redis.Client
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, rdb *redis.Client, db *gorm.DB) *GreeterService {
	return &GreeterService{
		uc:  uc,
		Db:  db,
		Rdb: rdb,
	}
}

func isValidMobile(mobile string) bool {
	pattern := `^1[3-9]\d{9}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(mobile)
}

func (s *GreeterService) SendSms(ctx context.Context, in *v1.SendSmsRequest) (*v1.SendSmsReply, error) {
	//判断手机号是否正确
	if !isValidMobile(in.Mobile) {
		return nil, fmt.Errorf("手机号格式不正确")
	}

	// 使用Redis限制发送频率
	set, err := s.Rdb.SetNX(ctx, "SendSmsLimit:"+in.Mobile, 1, 60*time.Second).Result()
	if err != nil {
		return nil, fmt.Errorf("redis错误: %v", err)
	}
	if !set {
		return nil, fmt.Errorf("短信发送过于频繁，请60秒后再试")
	}

	// 生成验证码
	code := rand.Intn(9000) + 1000

	// 将生成的验证码存储到redis缓存中，并设置过期时间5分钟
	err = s.Rdb.Set(ctx, "SendSms:"+in.Mobile, code, 5*time.Minute).Err()
	if err != nil {
		return nil, fmt.Errorf("redis错误: %v", err)
	}

	return &v1.SendSmsReply{
		Code: 200,
		Msg:  "发送验证码成功",
	}, nil
}

func (s *GreeterService) Register(_ context.Context, in *v1.RegisterRequest) (*v1.RegisterReply, error) {
	//判断手机号
	if !isValidMobile(in.Mobile) {
		return nil, fmt.Errorf("手机号格式不正确")
	}

	var user biz.User
	s.Db.Where("mobile=?", in.Mobile).Find(&user)
	// 判断用户是否存在，如果不存在直接注册

	if user.Id == 0 {
		newUser := biz.User{
			Mobile:   in.Mobile,
			UserName: "用户" + in.Mobile,
			NickName: "小明",
			Sex:      "男",
			Mileage:  "150",
		}

		if err := s.Db.Create(&newUser).Error; err != nil {
			return nil, fmt.Errorf("注册失败")
		}
	}
	// 接受验证码，判断验证码是否过期和验证码是否错误
	get := s.Rdb.Get(context.Background(), "SendSms:"+in.Mobile)
	fmt.Println(get)
	if get.Err() != nil {
		return nil, fmt.Errorf("验证码已过期")
	}

	if get.Val() != in.SendSmsCode {
		return nil, fmt.Errorf("验证码错误")
	}
	// 5.验证成功后，删除验证码
	s.Rdb.Del(context.Background(), "SendSms"+in.Mobile)

	token, _ := pkg.NewJWT("2211a").CreateToken(pkg.CustomClaims{
		ID: uint(user.Id),
	})

	return &v1.RegisterReply{
		Code:  200,
		Msg:   "登录成功",
		Token: token,
	}, nil
}
