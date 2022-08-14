// Package go_redis_client
// Time    : 2022/8/13 11:34
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package go_redis_client

import "github.com/go-redis/redis/v8"

type RedisClient struct {
	r *redis.Client
}

type RedisConf struct {
	// host:port address.
	Addr     string
	Username string
	Password string
	DB       int
}

func NewRedisClient(c *RedisConf) *RedisClient {
	return &RedisClient{
		r: redis.NewClient(&redis.Options{
			Addr:     c.Addr,
			Username: c.Username,
			Password: c.Password,
			DB:       c.DB,
		}),
	}
}
