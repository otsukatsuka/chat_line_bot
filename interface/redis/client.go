package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/otsukatsuka/chat_line_bot/config"
)

type Redis interface {
	GetConnection() redis.Conn
}

type redisClient struct {
	redisPool *redis.Pool
}

func (r redisClient) GetConnection() redis.Conn {
	return r.redisPool.Get()
}

func NewRedisClient() (Redis, error) {
	cfg, err := newRedisConfig()
	if err != nil {
		return nil, err
	}
	redisAddr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)
	pool := redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisAddr)
		},
		MaxConnLifetime: 10,
	}
	return &redisClient{redisPool: &pool}, nil
}
func newRedisConfig() (*config.RedisConfig, error) {
	cfg, err := config.NewRedisConfig()
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
