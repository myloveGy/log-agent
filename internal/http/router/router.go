package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log-agent/internal/config"
	"log-agent/internal/http/api"
	"log-agent/internal/http/middleware"
	"log-agent/internal/http/response"
	"log-agent/internal/util"
)

type Router struct {
	User       *api.User
	Database   *api.Database
	Guest      *api.Guest
	Config     *config.Config
	Middleware *middleware.Middleware
}

func (r *Router) Run(addr string) error {
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
					"errors":  e.Errors,
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
	app.Use(cors.New(cors.Config{
		AllowOrigins: r.Config.HttpConfig.AllowOrigins,
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"name":  r.Config.AppName,
			"start": r.Config.Start.Format("2006-01-02 15:04:05"),
			"date":  util.DateTime(),
			"ip":    c.IP(),
		})
	})

	// 游客可访问路由
	app.Post("/login", r.Guest.Login)
	app.Post("/register", r.Guest.Register)
	app.Get("/allow-register", r.Guest.AllowRegister)

	auth := app.Use(r.Middleware.Auth())
	auth.Post("/user/detail", r.User.Detail)
	auth.Post("/user/create", r.User.Create)
	auth.Post("/user/update", r.User.Update)
	auth.Post("/user/list", r.User.List)

	// 数据库操作
	auth.Post("/database/query", r.Database.Query)
	auth.Post("/database/collections", r.Database.Collections)

	return app
}
