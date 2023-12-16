package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/wegoteam/wepkg/config"
	"sync"
)

const (
	REDIS = "redis"
)

var (
	RedisCliet *redis.Client
	once       sync.Once
)

// init
// @Description: 初始化配置
func init() {
	once.Do(func() {
		initRedisConfig()
	})
}

// initRedisConfig
// @Description: 初始化Redis配置
func initRedisConfig() {
	var redisConfig = &config.Redis{}
	c := config.GetConfig()
	c.Load(REDIS, redisConfig)
	RedisCliet = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	_, redisErr := RedisCliet.Ping(context.Background()).Result()
	if redisErr != nil {
		fmt.Errorf("redis init error: %v", redisErr)
	}
}
