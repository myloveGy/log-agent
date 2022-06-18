package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id            primitive.ObjectID `bson:"_id" json:"id"`
	Username      string             `json:"username" bson:"username"`
	Password      string             `json:"password" bson:"password"`
	Status        string             `json:"status" bson:"status"`
	LastLoginTime string             `json:"last_login_time" bson:"last_login_time"`
	LastLoginIp   string             `json:"last_login_ip" bson:"last_login_ip"`
	CreatedAt     string             `json:"created_at" bson:"created_at"`
	UpdatedAt     string             `json:"updated_at" bson:"updated_at"`
}

func (*User) Collection() string {
	return "user"
}
