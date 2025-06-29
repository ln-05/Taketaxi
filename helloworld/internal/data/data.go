package data

import (
	"context"
	"helloworld/internal/biz"
	"helloworld/internal/conf"

	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewDB, NewRedis)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, rdb *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		// 关闭数据库连接
		if db != nil {
			sqlDB, err := db.DB()
			if err == nil {
				sqlDB.Close()
			}
		}
		// 关闭 Redis 连接
		if rdb != nil {
			rdb.Close()
		}
	}
	return &Data{
		db:  db,
		rdb: rdb,
	}, cleanup, nil
}

// NewDB 创建数据库连接
func NewDB(c *conf.Data, logger log.Logger) (*gorm.DB, error) {
	log.NewHelper(logger).Infof("connecting to database: %s", c.Database.Source)

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 获取底层的 sql.DB 对象来设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// 自动迁移表结构
	err = db.AutoMigrate(&biz.User{})
	if err != nil {
		log.NewHelper(logger).Warnf("auto migrate failed: %v", err)
	} else {
		log.NewHelper(logger).Info("auto migrate completed")
	}

	log.NewHelper(logger).Info("database connected successfully")
	return db, nil
}

// NewRedis 创建 Redis 连接
func NewRedis(c *conf.Data, logger log.Logger) (*redis.Client, error) {
	log.NewHelper(logger).Infof("connecting to redis: %s", c.Redis.Addr)

	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           0, // 使用默认数据库
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	})

	// 测试连接
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	log.NewHelper(logger).Info("redis连接成功")
	return rdb, nil
}

// GetDB 获取数据库连接
func (d *Data) GetDB() *gorm.DB {
	return d.db
}

// GetRedis 获取 Redis 连接
func (d *Data) GetRedis() *redis.Client {
	return d.rdb
}
