// Package go_redis_client
// Time    : 2022/8/13 11:34
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package go_redis_client

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

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

func (c *RedisClient) Close() error {
	return c.r.Close()
}

func (c *RedisClient) Flush(ctx context.Context) error {
	return c.r.FlushDB(ctx).Err()
}

func (c *RedisClient) Set(ctx context.Context, key string, value interface{}) error {
	return c.r.Set(ctx, key, value, 0).Err()
}

func (c *RedisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := c.r.Get(ctx, key).Result()
	if err != nil {
		return "", errors.Wrapf(err, "redis.Get %s", key)
	}
	return val, nil
}

func (c *RedisClient) HyperLogAdd(ctx context.Context, key string, els ...interface{}) (int64, error) {
	val, err := c.r.PFAdd(ctx, key, els...).Result()
	if err != nil {
		return 0, errors.Wrapf(err, "redis.PFAdd %s", key)
	}
	return val, nil
}

func (c *RedisClient) HyperLogCount(ctx context.Context, key string) (int64, error) {
	val, err := c.r.PFCount(ctx, key).Result()
	if err != nil {
		return 0, errors.Wrapf(err, "redis.PFCount %s", key)
	}
	return val, nil
}
