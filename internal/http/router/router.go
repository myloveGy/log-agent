package router

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log-agent/internal/http/api"
	"log-agent/internal/http/response"
	"log-agent/internal/util"
)

type Router struct {
	User     *api.User
	Database *api.Database
}

func (r *Router) Run(ctx context.Context, addr string) error {
	app := r.register()
	return app.Listen(addr)
}

func (r *Router) register() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if e, ok := err.(*response.Response); ok {
				return ctx.Status(e.Status).JSON(e)
			}

			if e, ok := err.(*util.ErrorResponse); ok {
				return ctx.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
					"items":   e.Items,
					"message": e.Error(),
					"code":    response.InvalidParameter,
				})
			}

			return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{
				"message": err.Error(),
				"code":    string(response.SystemError),
			})
		},
	})

	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"name": "log-agent",
			"date": util.DateTime(),
			"ip":   c.IP(),
		})
	})

	app.Post("/login", r.User.Login)
	app.Post("/user/create", r.User.Create)
	app.Post("/user/update", r.User.Update)

	app.Post("/database/query", r.Database.Query)

	return app
}