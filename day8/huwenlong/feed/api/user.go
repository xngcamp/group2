package api

import "github.com/globalsign/mgo/bson"

type User struct {
	UserId bson.ObjectId `json:"id" bson:"_id"`
	Nick string `json:"nick" bson:"_nick"`
	Sex int `json:"sex" bson:"_sex"`
	Email string `json:"email" bson:"_email"`
	Password string `json:"password" bson:"_password"`
}

func NewUser() *User {
	return &User{}
}
