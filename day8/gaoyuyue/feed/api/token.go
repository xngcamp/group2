package api

import "gopkg.in/mgo.v2/bson"

type Token struct {
	Id bson.ObjectId `json:"id"`
	UserId bson.ObjectId `json:"user_id"`
}

func NewToken() *Token {
	return &Token{}
}
