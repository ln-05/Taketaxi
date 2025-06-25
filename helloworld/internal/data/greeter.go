package data

import (
	"context"
	"time"

	"helloworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// Greeter 问候语模型
type Greeter struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	Hello     string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

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

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	greeter := &Greeter{
		Hello: g.Hello,
	}

	if err := r.data.Db.WithContext(ctx).Create(greeter).Error; err != nil {
		return nil, err
	}

	return &biz.Greeter{
		Hello: greeter.Hello,
	}, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	greeter := &Greeter{
		Hello: g.Hello,
	}

	if err := r.data.Db.WithContext(ctx).Save(greeter).Error; err != nil {
		return nil, err
	}

	return &biz.Greeter{
		Hello: greeter.Hello,
	}, nil
}

func (r *greeterRepo) FindByID(ctx context.Context, id int64) (*biz.Greeter, error) {
	var greeter Greeter
	if err := r.data.Db.WithContext(ctx).First(&greeter, id).Error; err != nil {
		return nil, err
	}

	return &biz.Greeter{
		Hello: greeter.Hello,
	}, nil
}

func (r *greeterRepo) ListByHello(ctx context.Context, hello string) ([]*biz.Greeter, error) {
	var greeters []Greeter
	if err := r.data.Db.WithContext(ctx).Where("hello LIKE ?", "%"+hello+"%").Find(&greeters).Error; err != nil {
		return nil, err
	}

	result := make([]*biz.Greeter, 0, len(greeters))
	for _, g := range greeters {
		result = append(result, &biz.Greeter{
			Hello: g.Hello,
		})
	}

	return result, nil
}

func (r *greeterRepo) ListAll(ctx context.Context) ([]*biz.Greeter, error) {
	var greeters []Greeter
	if err := r.data.Db.WithContext(ctx).Find(&greeters).Error; err != nil {
		return nil, err
	}

	result := make([]*biz.Greeter, 0, len(greeters))
	for _, g := range greeters {
		result = append(result, &biz.Greeter{
			Hello: g.Hello,
		})
	}

	return result, nil
}
