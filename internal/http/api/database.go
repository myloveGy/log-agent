package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"log-agent/internal/http/request"
	"log-agent/internal/http/response"
	"log-agent/internal/repo"
	"log-agent/internal/util"
)

type Database struct {
	UserRepo *repo.UserRepo
}

func (d *Database) Collections(c *fiber.Ctx) error {
	items, err := d.UserRepo.Database.ListCollectionNames(c.Context(), bson.M{"name": bson.M{"$ne": "user"}})
	if err != nil {
		return response.NewSystemError(err)
	}

	return c.JSON(map[string]interface{}{
		"items": items,
	})
}

func (d *Database) Query(c *fiber.Ctx) error {

	// 解析参数
	param := &request.DatabaseQueryRequest{}
	if err := c.BodyParser(param); err != nil {
		return response.NewInvalidParameter(err)
	}

	// 验证参数
	if err := util.Validate(param); err != nil {
		return err
	}

	// 排序方式
	sortValue := -1
	if param.Order == "asc" {
		sortValue = 1
	}

	// 时间段查询
	if param.StartTime == "" {
		param.StartTime = fmt.Sprintf("%s 00:00:00", util.Date())
	}

	if param.EndTime == "" {
		param.EndTime = fmt.Sprintf("%s 23:59:59", util.Date())
	}

	param.Query["datetime"] = bson.M{
		"$gte": param.StartTime,
		"$lte": param.EndTime,
	}

	// 统计数据
	results := make([]map[string]interface{}, 0)
	pagination, err := d.UserRepo.Pagination(c.Context(), &results, &repo.Builder{
		Collection: param.Collection,
		Query:      param.Query,
		Page:       param.Page,
		PageSize:   param.PageSize,
		Order:      "datetime",
		Sort:       sortValue,
	})
	if err != nil {
		return response.NewSystemError(err)
	}

	// 没有数据、提前返回
	return c.JSON(pagination)
}
