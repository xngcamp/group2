package api

import "github.com/globalsign/mgo/bson"

type User struct {
	Nick string `json:"nick"`	//用户姓名
	Sex int `json:"sex"`	//用户性别（1|2）
	Email string `json:"email"` 	//用户邮箱，做登录用，并且检测是否已经注册
	Password string `json:"password"` 	//用户的密码（后期加密）
	Token bson.ObjectId `json:"token" bson:"_id"`
}

func NewUser() *User {
	return &User{}
}
