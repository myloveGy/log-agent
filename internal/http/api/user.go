package api

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (u *User) Login(c *fiber.Ctx) error {
	param := &request.UserRequest{}
	if err := c.BodyParser(param); err != nil {
		return response.NewInvalidParameter(err)
	}

	if err := util.Validate(param); err != nil {
		return err
	}

	item, err := u.UserRepo.GetUser(c.Context(), param.Username)
	if err != nil {
		return response.NewInvalidParameter(errors.New("用户名或者密码错误"))
	}

	// 验证密码
	if err := util.ValidatePassword(param.Password, item.Password); err != nil {
		return response.NewInvalidParameter(errors.New("用户名或者密码错误"))
	}

	// 修改登录信息
	if err := u.UserRepo.Update(c.Context(), &model.User{
		Id:            item.Id,
		LastLoginTime: util.DateTime(),
		LastLoginIp:   c.IP(),
		UpdatedAt:     util.DateTime(),
	}); err != nil {
		return response.NewSystemError(err)
	}

	var token string
	if token, err = util.GenerateToken(map[string]string{
		"username": param.Username,
	}, u.Config.Jwt.Secret, time.Duration(u.Config.Jwt.ExpiresTime)*time.Second); err != nil {
		return response.NewSystemError(err)
	}

	return c.JSON(map[string]interface{}{
		"token":           token,
		"username":        item.Username,
		"last_login_time": item.LastLoginTime,
		"last_login_ip":   item.LastLoginIp,
		"created_at":      item.CreatedAt,
	})
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
	total, err := u.UserRepo.Database.Collection("user").CountDocuments(c.Context(), param.Query)
	if err != nil {
		return response.NewSystemError(err)
	}

	results := make([]map[string]interface{}, 0)
	pagination := &response.Pagination{
		Items:    results,
		Page:     param.Page,
		PageSize: param.PageSize,
		Total:    total,
	}

	// 没有数据、提前返回
	if total == 0 {
		return c.JSON(pagination)
	}

	// 分页查询数据
	sort := bson.D{{"created_at", 1}}
	findOptions := options.Find()
	findOptions.SetSort(sort)
	findOptions.SetLimit(param.PageSize)
	findOptions.SetSkip(param.PageSize * (param.Page - 1))

	var cursor *mongo.Cursor
	if cursor, err = u.UserRepo.Database.Collection("user").Find(c.Context(), param.Query, findOptions); err != nil {
		return response.NewSystemError(err)
	}

	defer cursor.Close(c.Context())

	for cursor.Next(c.Context()) {
		item := make(map[string]interface{})
		if err := cursor.Decode(item); err != nil {
			continue
		}

		results = append(results, item)
	}

	pagination.Items = results
	return c.JSON(pagination)
}
