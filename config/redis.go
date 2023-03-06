package config

import "github.com/kelseyhightower/envconfig"

type RedisConfig struct {
	RedisPort string `envconfig:"REDIS_PORT" default:"6379"`
	RedisHost string `envconfig:"REDIS_HOST" default:""`
}

func NewRedisConfig() (RedisConfig, error) {
	var redisConfig RedisConfig
	err := envconfig.Process("", &redisConfig)
	return redisConfig, err
}
