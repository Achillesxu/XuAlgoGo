// Package go_redis_client
// Time    : 2022/8/13 12:13
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package go_redis_client

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewRedisClient(t *testing.T) {
	rc := RedisConf{
		Addr:     "localhost:6379",
		Username: "",
		Password: "",
		DB:       0,
	}
	req := require.New(t)
	ctx := context.Background()
	rClient := NewRedisClient(&rc)
	err := rClient.r.Ping(ctx).Err()
	req.NoError(err)

}
