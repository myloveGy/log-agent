package util

import (
	"fmt"
	"testing"

	"log-agent/internal/model"
)

func TestStructNotZero2BsonM(t *testing.T) {
	user := &model.User{Username: DateTime()}
	bsonM := StructNotZero2BsonM(user)
	fmt.Println(bsonM)

	bsonM = StructNotZero2BsonM(nil)
	fmt.Println(bsonM)
}
