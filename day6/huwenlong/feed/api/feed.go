package api

import "github.com/globalsign/mgo/bson"

//创建feed结构体
type Feed struct {
	Txt string `json:"txt" bson:"_txt"`
	Id bson.ObjectId `json:"id" bson:"_id"`
}

//暴露feed结构体 生成feed
func NewFeed()  *Feed{
	return &Feed{}
}