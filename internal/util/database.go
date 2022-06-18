package util

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

func StructNotZero2BsonM(value interface{}) bson.M {
	fields := bson.M{}

	// 值
	valueOf := reflect.ValueOf(value)
	valueType := reflect.TypeOf(value)
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
		valueType = valueType.Elem()
	}

	if valueOf.Kind() != reflect.Struct {
		return fields
	}

	for i := 0; i < valueType.NumField(); i++ {
		if !valueOf.Field(i).IsZero() {
			fields[valueType.Field(i).Tag.Get("bson")] = valueOf.Field(i).Interface()
		}
	}

	return fields
}
