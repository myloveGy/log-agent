package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"log-agent/internal/config"
	"log-agent/internal/http/response"
	"log-agent/internal/repo"
	"log-agent/internal/util"
)

type Auth func() fiber.Handler

func NewAuth(conf *config.Config, userRepo *repo.UserRepo) Auth {
	return func() fiber.Handler {
		return func(c *fiber.Ctx) error {
			token := c.Get("Authorization")
			if token == "" {
				token = c.Query("token")
			}

			token = strings.TrimSpace(strings.TrimLeft(token, "Bearer"))
			if token == "" {
				return response.NewUnauthenticated("Bearer Token is empty")
			}

			item := make(map[string]string)
			if _, err := util.ParseToken(&item, token, conf.Jwt.Secret); err != nil {
				return response.NewUnauthenticated(fmt.Sprintf("Bearer Token error: %s", err))
			}

			user, err := userRepo.GetUser(c.Context(), item["username"])
			if err != nil {
				return response.NewUnauthenticated(fmt.Sprintf("Bearer Token error: user not found %s", err))
			}

			// 存储用户信息
			c.Locals("user", user)

			return c.Next()
		}
	}
}
