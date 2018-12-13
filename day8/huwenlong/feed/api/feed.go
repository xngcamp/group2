package api

import "github.com/globalsign/mgo/bson"

//创建feed结构体
type Feed struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Txt string `json:"txt" bson:"_txt"`

}

//暴露feed结构体 生成feed
func NewFeed()  *Feed{
	return &Feed{}
}