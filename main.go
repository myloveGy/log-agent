package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/sync/errgroup"
	"log-agent/config"
	"log-agent/connect"
	"log-agent/tail"
)

var (
	conf        = &config.Config{}
	client      *mongo.Client
	redisClient *redis.Client
	err         error
)

func main() {

	// 第一步：加载配置文件
	conf, err = config.Load("./.ini")
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eg, _ := errgroup.WithContext(ctx)

	// 第二步：连接mongo
	client, err = connect.MongoConnect(ctx, conf.MongoDB)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("MongoDB Closed Error: %v\n", err)
		}

		log.Println("Connection to MongoDB closed.")
	}()

	// 连接redis:
	redisClient, err = connect.RedisConnect(ctx, conf.Redis)
	if err != nil {
		log.Fatalln(err)
	}

	// 注册mongo处理
	tail.RegisterDriver("mongo", func(data interface{}, tailConfig *tail.Config) error {
		collection := client.Database(conf.MongoDB.DBName).Collection(tailConfig.Name)
		_, err := collection.InsertOne(ctx, data)
		return err
	})

	// 注册redis处理
	tail.RegisterDriver("redis", func(data interface{}, conf *tail.Config) error {
		if conf.Type != "text" {
			str, err := json.Marshal(data)
			if err != nil {
				return err
			}

			return redisClient.LPush(ctx, conf.Name, string(str)).Err()
		}

		return redisClient.LPush(ctx, conf.Name, data).Err()
	})

	// 第三步：处理配置文件
	group := &sync.WaitGroup{}
	for _, v := range conf.Handles {
		handle, err := tail.NewTailHandle(v)
		if err != nil {
			log.Fatalln(err)
		}

		group.Add(1)
		log.Printf("Tail Handle Start: %s\n", v.Filename)
		handle.Run(ctx, group)
	}

	eg.Go(func() error {
		group.Wait()
		return nil
	})

	// 最后监听配置退出
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c:
			// 监听的退出信号
			log.Printf("Service Exit")
			cancel()
			return nil
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatalf("eg error: %s", err)
	}
}
