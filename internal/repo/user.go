package repo

import (
	"context"
	"errors"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log-agent/internal/model"
	"log-agent/internal/util"
)

type UserRepo struct {
	Database *mongo.Database
}

func (u *UserRepo) Count(ctx context.Context, user *model.User) (int64, error) {
	return u.Database.Collection(user.Collection()).CountDocuments(ctx, util.StructNotZero2BsonM(user))
}

func (u *UserRepo) GetUser(ctx context.Context, username string) (*model.User, error) {
	item := &model.User{}
	if err := u.Database.Collection(item.Collection()).FindOne(ctx, bson.M{
		"username": username,
	}).Decode(item); err != nil {
		return nil, err
	}

	return item, nil
}

func (u *UserRepo) Create(ctx context.Context, user *model.User) error {
	if _, err := u.Database.Collection(user.Collection()).InsertOne(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) Update(ctx context.Context, user *model.User) error {
	if _, err := u.Database.Collection(user.Collection()).UpdateOne(ctx, bson.M{
		"_id": user.Id,
	}, bson.D{
		{"$set", util.StructNotZero2BsonM(user)},
	}); err != nil {
		return err
	}

	return nil
}

type Pagination struct {
	Items    interface{} `json:"items"`
	Total    int64       `json:"total"`
	Page     int64       `json:"page"`
	PageSize int64       `json:"page_size"`
}

type Builder struct {
	Collection string      `json:"collection"`
	Query      interface{} `json:"query"`
	Page       int64       `json:"page"`
	PageSize   int64       `json:"page_size"`
	Order      string      `json:"order"`
	Sort       int         `json:"sort"`
}

func (u *UserRepo) Pagination(ctx context.Context, items interface{}, builder *Builder) (*Pagination, error) {
	// 统计数据
	var (
		pagination = &Pagination{
			Items:    items,
			Page:     builder.Page,
			PageSize: builder.PageSize,
		}
		err error
	)

	// 分页最大1000
	if pagination.PageSize <= 0 {
		pagination.PageSize = 10
	} else if pagination.PageSize > 1000 {
		pagination.PageSize = 1000
	}

	// 最少第一页
	if pagination.Page <= 0 {
		pagination.Page = 1
	}

	if pagination.Total, err = u.Database.Collection(builder.Collection).CountDocuments(ctx, builder.Query); err != nil {
		return nil, err
	}

	// 没有数据、提前返回
	if pagination.Total == 0 {
		return pagination, nil
	}

	// 分页查询数据
	findOptions := options.Find()
	if builder.Order != "" {
		findOptions.SetSort(bson.M{builder.Order: builder.Sort})
	}

	findOptions.SetLimit(builder.PageSize)
	findOptions.SetSkip(builder.PageSize * (builder.Page - 1))
	var cursor *mongo.Cursor
	if cursor, err = u.Database.Collection(builder.Collection).Find(ctx, builder.Query, findOptions); err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	value := reflect.ValueOf(items)
	// json.Unmarshal returns errors for these
	if value.Kind() != reflect.Ptr {
		return nil, errors.New("must pass a pointer, not a value, to StructScan destination")
	}

	if value.IsNil() {
		return nil, errors.New("nil pointer passed to StructScan destination")
	}

	valueType := value.Type()
	if valueType.Kind() == reflect.Ptr {
		valueType = valueType.Elem()
	}

	if valueType.Kind() != reflect.Slice {
		return nil, errors.New("must pass a pointer to a slice, not a value, to StructScan destination")
	}

	isPtr := valueType.Elem().Kind() == reflect.Ptr
	base := valueType.Elem()
	if base.Kind() == reflect.Ptr {
		base = base.Elem()
	}

	direct := reflect.Indirect(value)
	for cursor.Next(ctx) {
		vp := reflect.New(base)
		if err := cursor.Decode(vp.Interface()); err != nil {
			continue
		}

		// append
		if isPtr {
			direct.Set(reflect.Append(direct, vp))
		} else {
			direct.Set(reflect.Append(direct, reflect.Indirect(vp)))
		}
	}

	return pagination, nil
}
