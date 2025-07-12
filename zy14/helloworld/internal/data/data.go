package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"helloworld/internal/conf"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewRedis, NewMysql)

//var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewGreeterRepo, NewSmsRepo, , NewDB, NewMinio)

// Data .
type Data struct {
	DB  *gorm.DB
	Rdb *redis.Client
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, rdb *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		DB:  db,
		Rdb: rdb,
	}, cleanup, nil
}
func NewRedis(c *conf.Data, logger log.Logger) (*redis.Client, func(), error) {
	log := log.NewHelper(logger)

	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           0,
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	})

	// 测试Redis连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Errorf("failed to connect to redis: %v", err)
		return nil, nil, err
	}

	return rdb, func() {
		rdb.Close()
	}, nil

}

// NewDB 创建数据库连接
func NewMysql(c *conf.Data, logger log.Logger) (*gorm.DB, func(), error) {
	log := log.NewHelper(logger)
	// 初始化数据库连接
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		log.Errorf("failed opening connection to mysql: %v", err)
		return nil, nil, err
	}

	// 获取底层的sql.DB对象以设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Errorf("failed to get sql.DB: %v", err)
		return nil, nil, err
	}
	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, func() {
		sqlDB.Close()
	}, nil
}
