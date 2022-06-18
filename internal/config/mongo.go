package config

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Database string `ini:"database"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

func NewMongoConfig(config *Config) *MongoConfig {
	return config.Mongo
}

func NewMongoDatabase(conf *MongoConfig) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 设置了登录的账号和密码
	loginAuth := ""
	if conf.Username != "" && conf.Password != "" {
		loginAuth = fmt.Sprintf("%s:%s@", conf.Username, conf.Password)
	}

	// 连接mongo db
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s%s:%d", loginAuth, conf.Host, conf.Port))
	clientOptions.SetMaxPoolSize(10)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	return client.Database(conf.Database)
}
