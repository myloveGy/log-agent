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
	// 设置了登录的账号和密码
	loginAuth := ""
	if conf.Username != "" && conf.Password != "" {
		loginAuth = fmt.Sprintf("%s:%s@", conf.Username, conf.Password)
	}

	// 连接mongo db
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s%s:%d", loginAuth, conf.Host, conf.Port))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	if conf.Username != "" && conf.Password != "" {

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
