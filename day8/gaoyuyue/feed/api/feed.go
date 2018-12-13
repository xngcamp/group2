package api

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Feed struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	UserId bson.ObjectId `json:"user_id"`
	CreateTime time.Time `json:"create_time"`
}

func NewFeed() *Feed {
	return &Feed{}
}