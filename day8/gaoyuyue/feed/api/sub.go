package api

import "gopkg.in/mgo.v2/bson"

type Sub struct {
	Id bson.ObjectId
	Followers []bson.ObjectId
}

func NewSub() *Sub {
	return &Sub{}
}