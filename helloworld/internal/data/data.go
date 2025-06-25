package data

import (
	"fmt"
	"helloworld/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewRedisClient)

// Data .
type Data struct {
	// TODO wrapped database client
	Db  *gorm.DB
	Rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	// 1. 初始化 MySQL
	db, err := gorm.Open(mysql.Open(c.GetDatabase().GetSource()), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("mysql连接成功")

	// 2. 初始化 Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.GetRedis().GetAddr(),
		Password: c.GetRedis().GetPassword(),
		// 你可以根据需要加上 DB、PoolSize、Timeout 等参数
	})

	fmt.Println("redis连接成功")

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		sqlDB, _ := db.DB()
		sqlDB.Close()
		rdb.Close()
	}
	return &Data{Db: db, Rdb: rdb}, cleanup, nil
}

// NewRedisClient 提供Redis客户端
func NewRedisClient(c *conf.Data) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.GetRedis().GetAddr(),
		Password: c.GetRedis().GetPassword(),
		// 你可以根据需要加上 DB、PoolSize、Timeout 等参数
	})
}
