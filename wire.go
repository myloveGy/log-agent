//go:build wireinject
// +build wireinject

package main

import (
	"github.com/urfave/cli/v2"
	"log-agent/internal/cmd"
	"log-agent/internal/config"
	"log-agent/internal/http/api"
	"log-agent/internal/http/middleware"
	"log-agent/internal/http/router"
	"log-agent/internal/repo"

	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	// 配置信息
	config.NewMongoConfig,
	config.NewMongoDatabase,
	cmd.ProviderSet,

	// DB处理
	wire.Struct(new(repo.UserRepo), "*"),

	// 路由API
	middleware.NewAuth,
	wire.Struct(new(middleware.Middleware), "*"),
	wire.Struct(new(api.User), "*"),
	wire.Struct(new(api.Database), "*"),
	wire.Struct(new(router.Router), "*"),
)

func bootstrap(conf *config.Config) (cli.Commands, func()) {
	panic(wire.Build(providerSet))
}
