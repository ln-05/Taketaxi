package data

import (
	"context"
	"encoding/json"
	"fmt"

	"helloworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 示例数据模型
type User struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"size:255;not null"`
	Age  int    `gorm:"not null"`
}

// 示例：从数据库获取用户
func (r *greeterRepo) GetUser(ctx context.Context, id uint) (*User, error) {
	var user User
	err := r.data.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 示例：创建用户
func (r *greeterRepo) CreateUser(ctx context.Context, user *User) error {
	return r.data.db.WithContext(ctx).Create(user).Error
}

// 示例：从 Redis 缓存获取数据
func (r *greeterRepo) GetFromCache(ctx context.Context, key string) (string, error) {
	return r.data.rdb.Get(ctx, key).Result()
}

// 示例：设置 Redis 缓存
func (r *greeterRepo) SetCache(ctx context.Context, key, value string) error {
	return r.data.rdb.Set(ctx, key, value, 0).Err()
}

// 示例：使用 Redis 缓存用户信息
func (r *greeterRepo) GetUserWithCache(ctx context.Context, id uint) (*User, error) {
	// 先从缓存获取
	cacheKey := fmt.Sprintf("user:%d", id)
	cachedData, err := r.GetFromCache(ctx, cacheKey)
	if err == nil {
		// 缓存命中，解析数据
		var user User
		if err := json.Unmarshal([]byte(cachedData), &user); err == nil {
			return &user, nil
		}
	}

	// 缓存未命中，从数据库获取
	user, err := r.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	// 将数据存入缓存
	if userData, err := json.Marshal(user); err == nil {
		r.SetCache(ctx, cacheKey, string(userData))
	}

	return user, nil
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}
