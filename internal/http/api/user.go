package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log-agent/internal/config"
	"log-agent/internal/http/request"
	"log-agent/internal/http/response"
	"log-agent/internal/model"
	"log-agent/internal/repo"
	"log-agent/internal/util"
)

type User struct {
	UserRepo *repo.UserRepo
	Config   *config.Config
}

func (u *User) Detail(c *fiber.Ctx) error {
	return c.JSON(c.Locals("user"))
}

func (u *User) Create(c *fiber.Ctx) error {
	param := &request.UserRequest{}
	if err := c.BodyParser(param); err != nil {
		return response.NewInvalidParameter(err)
	}

	if err := util.Validate(param); err != nil {
		return err
	}

	if _, err := u.UserRepo.GetUser(c.Context(), param.Username); !errors.Is(err, mongo.ErrNoDocuments) {
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

	if err := u.UserRepo.Create(c.Context(), user); err != nil {
		return response.NewSystemError(err)
	}

	return c.JSON(map[string]string{
		"username": param.Username,
	})
}

func (u *User) Update(c *fiber.Ctx) error {
	param := &request.UserUpdateRequest{}
	if err := c.BodyParser(param); err != nil {
		return response.NewInvalidParameter(err)
	}

	if err := util.Validate(param); err != nil {
		return err
	}

	item, err := u.UserRepo.GetUser(c.Context(), param.Username)
	if err != nil {
		return response.NewSystemError(err)
	}

	updateModel := &model.User{
		Id:        item.Id,
		Username:  param.Username,
		UpdatedAt: util.DateTime(),
	}

	if param.Password != "" {
		if password, err := util.GeneratePassword(param.Password); err != nil {
			return response.NewInvalidParameter(errors.New("生成密码错误"))
		} else {
			updateModel.Password = string(password)
		}
	}

	if param.Status != "" {
		updateModel.Status = param.Status
	}

	if err := u.UserRepo.Update(c.Context(), updateModel); err != nil {
		return response.NewSystemError(err)
	}

	return c.JSON(map[string]string{
		"username": param.Username,
	})
}

func (u *User) List(c *fiber.Ctx) error {

	// 解析参数
	param := &request.UserQueryRequest{}
	if err := c.BodyParser(param); err != nil {
		return response.NewInvalidParameter(err)
	}

	// 验证参数
	if err := util.Validate(param); err != nil {
		return err
	}

	// 统计数据
	results := make([]map[string]interface{}, 0)
	pagination, err := u.UserRepo.Pagination(c.Context(), &results, &repo.Builder{
		Collection: "user",
		Query:      param.Query,
		Page:       param.Page,
		PageSize:   param.PageSize,
		Sort:       1,
		Order:      "created_at",
	})
	if err != nil {
		return response.NewSystemError(err)
	}

	return c.JSON(pagination)
}
