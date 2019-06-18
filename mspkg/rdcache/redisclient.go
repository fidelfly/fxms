package rdcache

import (
	"fmt"

	"github.com/go-redis/redis"
)

var Redisdb *redis.Client

func InitRedis(cfg *Config) {
	Redisdb = NewRedisClient(cfg)
	return
}

func NewRedisClient(cfg *Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       0,
	})
}
