package connect

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log-agent/config"
)

func MongoConnect(ctx context.Context, conf config.MongoDBConfig) (*mongo.Client, error) {
	// 连接mongo db
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", conf.Host, conf.Port))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}

func RedisConnect(ctx context.Context, conf config.RedisConfig) (*redis.Client, error) {
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
