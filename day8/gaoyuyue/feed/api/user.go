package api

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Nick string `json:"nick"`
	Password string `json:"password"`
	Sex int `json:"sex"`
	Email string `json:"email"`
}

func NewUser() *User {
	return &User{}
}