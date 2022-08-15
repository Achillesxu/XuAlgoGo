// Package go_redis_client
// Time    : 2022/8/13 12:13
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
// testify suit setupsuit, teardownsuit, setup, teardown, test
package go_redis_client

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RedisTestSuit struct {
	r *RedisClient
	suite.Suite
}

func (s *RedisTestSuit) SetupSuite() {
	rc := RedisConf{
		Addr:     "localhost:6379",
		Username: "",
		Password: "",
		DB:       0,
	}
	s.r = NewRedisClient(&rc)
}

func (s *RedisTestSuit) TearDownSuite() {
	// ctx := context.Background()
	// _ = s.r.Flush(ctx)
	_ = s.r.Close()
}

// TestRedisFunc is all redis func test entrance
func TestRedisFunc(t *testing.T) {
	suite.Run(t, new(RedisTestSuit))
}

func (s *RedisTestSuit) TestRedisHyperLogAdd() {
	ctx := context.Background()
	key := "test_key"
	cnt := "h"
	res, err := s.r.HyperLogAdd(ctx, key, cnt)
	s.NoError(err)
	s.Equal(int64(1), res)
}

func (s *RedisTestSuit) TestRedisHyperLogCount() {
	ctx := context.Background()
	key := "test_key"
	cnt := "h"
	_, _ = s.r.HyperLogAdd(ctx, key, cnt)
	res, err := s.r.HyperLogCount(ctx, key)
	s.NoError(err)
	s.Equal(int64(2), res)
}
