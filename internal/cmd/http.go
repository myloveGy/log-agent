package cmd

import (
	"github.com/urfave/cli/v2"
	"log-agent/internal/http/router"
)

type HttpCmd *cli.Command

func NewHttpCmd(r *router.Router) HttpCmd {
	return &cli.Command{
		Name:  "http",
		Usage: "http command eg: ./app http --addr=:8080",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "addr",
				Usage:    "--addr=:8080",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			addr := c.String("addr")
			return r.Run(addr)
		},
	}
}
