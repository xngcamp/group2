package api

import "gopkg.in/mgo.v2/bson"

type Fans struct {
	Id bson.ObjectId
	Following []bson.ObjectId
}

func NewFans() *Fans {
	return &Fans{}
}