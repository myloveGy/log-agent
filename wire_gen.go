// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/urfave/cli/v2"
	"log-agent/internal/cmd"
	"log-agent/internal/config"
	"log-agent/internal/http/api"
	"log-agent/internal/http/router"
	"log-agent/internal/repo"
)

// Injectors from wire.go:

func bootstrap(conf *config.Config) (cli.Commands, func()) {
	mongoConfig := config.NewMongoConfig(conf)
	database := config.NewMongoDatabase(mongoConfig)
	userRepo := &repo.UserRepo{
		Database: database,
	}
	user := &api.User{
		UserRepo: userRepo,
	}
	apiDatabase := &api.Database{
		Database: database,
	}
	routerRouter := &router.Router{
		User:     user,
		Database: apiDatabase,
	}
	httpCmd := cmd.NewHttpCmd(routerRouter)
	tailCmd := cmd.NewTailCmd(conf)
	commands := &cmd.Commands{
		HttpCmd: httpCmd,
		TailCmd: tailCmd,
	}
	cliCommands := cmd.NewCommands(commands)
	return cliCommands, func() {
	}
}

// wire.go:

var providerSet = wire.NewSet(config.NewMongoConfig, config.NewMongoDatabase, cmd.ProviderSet, wire.Struct(new(repo.UserRepo), "*"), wire.Struct(new(api.User), "*"), wire.Struct(new(api.Database), "*"), wire.Struct(new(router.Router), "*"))
