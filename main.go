package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"log-agent/internal/config"
)

func main() {

	// 第一步：加载配置文件
	conf, err := config.Load("./config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	commands, cleanup := bootstrap(conf)
	defer cleanup()

	app := &cli.App{
		Name:     "log-agent",
		Version:  "0.0.1",
		Commands: commands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
