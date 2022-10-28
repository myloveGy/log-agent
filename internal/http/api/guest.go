package api

import (
	"errors"
	"strings"
	"time"

	"log-agent/internal/config"
	"log-agent/internal/http/request"
	"log-agent/internal/http/response"
	"log-agent/internal/model"
	"log-agent/internal/repo"
	"log-agent/internal/util"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Guest 游客接口
type Guest struct {
	UserRepo *repo.UserRepo
	Config   *config.Config
}

func (g *Guest) Login(c *fiber.Ctx) error {
	param := &request.UserRequest{}
	if err := c.BodyParser(param); err != nil {
		return response.NewInvalidParameter(err)
	}

	if err := util.Validate(param); err != nil {
		return err
	}

	item, err := g.UserRepo.GetUser(c.Context(), param.Username)
	if err != nil {
		return response.NewInvalidParameter(errors.New("用户名或者密码错误"))
	}

	// 验证密码
	if err := util.ValidatePassword(param.Password, item.Password); err != nil {
		return response.NewInvalidParameter(errors.New("用户名或者密码错误"))
	}

	// 验证状态
	if item.Status != model.StatusOnline {
		return response.NewInvalidParameter("用户已经停用")
	}

	// 修改登录信息
	if err := g.UserRepo.Update(c.Context(), &model.User{
		Id:            item.Id,
		LastLoginTime: util.DateTime(),
		LastLoginIp:   ClientIp(c),
		UpdatedAt:     util.DateTime(),
	}); err != nil {
		return response.NewSystemError(err)
	}

	var token string
	if token, err = g.token(item.Username); err != nil {
		return response.NewSystemError(err)
	}

	return c.JSON(&response.UserResponse{
		Token:         token,
		Username:      item.Username,
		LastLoginTime: item.LastLoginTime,
		LastLoginIp:   item.LastLoginIp,
		CreatedAt:     item.CreatedAt,
	})
}

func (g *Guest) Register(c *fiber.Ctx) error {

	// 判断是否已经存在
	total, err := g.UserRepo.Count(c.Context(), &model.User{})
	if err != nil {
		return response.NewSystemError(err)
	}

	// 不允许
	if total > 0 && !g.Config.HttpConfig.AllowRegister {
		return response.NewInvalidParameter("不允许注册")
	}

	param := &request.UserRequest{}
	if err := c.BodyParser(param); err != nil {
		return response.NewInvalidParameter(err)
	}

	if err := util.Validate(param); err != nil {
		return err
	}

	if _, err := g.UserRepo.GetUser(c.Context(), param.Username); !errors.Is(err, mongo.ErrNoDocuments) {
		return response.NewInvalidParameter(errors.New("用户名称已经存在"))
	}

	password, err := util.GeneratePassword(param.Password)
	if err != nil {
		return response.NewInvalidParameter(errors.New("生成密码错误"))
	}

	user := &model.User{
		Id:            primitive.NewObjectID(),
		Username:      param.Username,
		Password:      string(password),
		Status:        "Y",
		LastLoginTime: util.DateTime(),
		CreatedAt:     util.DateTime(),
		UpdatedAt:     util.DateTime(),
	}

	if err := g.UserRepo.Create(c.Context(), user); err != nil {
		return response.NewSystemError(err)
	}

	var token string
	if token, err = g.token(user.Username); err != nil {
		return response.NewSystemError(err)
	}

	return c.JSON(&response.UserResponse{
		Token:         token,
		Username:      user.Username,
		LastLoginTime: user.LastLoginTime,
		LastLoginIp:   user.LastLoginIp,
		CreatedAt:     user.CreatedAt,
	})
}

func (g *Guest) AllowRegister(c *fiber.Ctx) error {

	total, err := g.UserRepo.Count(c.Context(), &model.User{})
	if err != nil {
		return response.NewSystemError(err)
	}

	allow := model.StatusOffline
	if total == 0 || g.Config.HttpConfig.AllowRegister {
		allow = model.StatusOnline
	}

	return c.JSON(map[string]interface{}{
		"allow": allow,
		"total": total,
	})
}

func (g *Guest) token(username string) (string, error) {
	return util.GenerateToken(map[string]string{
		"username": username,
	}, g.Config.Jwt.Secret, time.Duration(g.Config.Jwt.ExpiresTime)*time.Second)
}

func ClientIp(c *fiber.Ctx) string {
	clientIp := string(c.Request().Header.Peek("X-Forwarded-For"))
	if index := strings.IndexByte(clientIp, ','); index >= 0 {
		clientIp = clientIp[0:index]
	}

	clientIp = strings.TrimSpace(clientIp)
	if len(clientIp) > 0 {
		return clientIp
	}

	clientIp = strings.TrimSpace(string(c.Request().Header.Peek("X-Real-Ip")))
	if len(clientIp) > 0 {
		return clientIp
	}

	return c.IP()
}
