package cmd

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/sync/errgroup"
	"log-agent/internal/config"
	"log-agent/internal/tail"
)

type TailCmd *cli.Command

func NewTailCmd(conf *config.Config, database *mongo.Database) TailCmd {
	return &cli.Command{
		Name:  "tail",
		Usage: "tail command eg: ./app tail",
		Action: func(cliCtx *cli.Context) error {
			ctx, cancel := context.WithCancel(cliCtx.Context)
			defer cancel()

			eg, _ := errgroup.WithContext(ctx)

			// 注册mongo处理
			tail.RegisterDriver("mongo", func(data interface{}, tailConfig *tail.Config) error {
				collection := database.Collection(tailConfig.Name)
				_, err := collection.InsertOne(ctx, data)
				return err
			})

			// 第三步：处理配置文件
			group := &sync.WaitGroup{}
			for _, v := range conf.Handler {
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

			return nil
		},
	}
}
