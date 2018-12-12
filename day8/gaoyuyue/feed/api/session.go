package api

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Session struct {
	Id bson.ObjectId `bson:"_id"`
	UserId bson.ObjectId
	Timeout time.Time
}

func NewSession() *Session {
	return &Session{}
}

