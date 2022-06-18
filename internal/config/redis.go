package config

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	DB       int    `ini:"db"`
}

func RedisConnect(ctx context.Context, conf *RedisConfig) (*redis.Client, error) {
	// 建立连接
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.DB,
	})

	// 检测心跳
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
