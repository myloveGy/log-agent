package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log-agent/internal/http/request"
	"log-agent/internal/http/response"
	"log-agent/internal/util"
)

type Database struct {
	Database *mongo.Database
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
	sortValue := 1
	if param.Order == "asc" {
		sortValue = -1
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
	total, err := d.Database.Collection(param.Collection).CountDocuments(c.Context(), param.Query)
	if err != nil {
		return response.NewSystemError(err)
	}

	// 没有数据、提前返回
	if total == 0 {
		return c.JSON(map[string]interface{}{
			"items":     results,
			"query":     param.Query,
			"total":     total,
			"page":      param.Page,
			"page_size": param.PageSize,
		})
	}

	// 分页查询数据
	sort := bson.D{{"dateline", sortValue}}
	findOptions := options.Find()
	findOptions.SetSort(sort)
	findOptions.SetLimit(param.PageSize)
	findOptions.SetSkip(param.PageSize * (param.Page - 1))

	var cursor *mongo.Cursor
	if cursor, err = d.Database.Collection(param.Collection).Find(c.Context(), param.Query, findOptions); err != nil {
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

	return c.JSON(map[string]interface{}{
		"items":     results,
		"total":     total,
		"page":      param.Page,
		"page_size": param.PageSize,
	})
}
