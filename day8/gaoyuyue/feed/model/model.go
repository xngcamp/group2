/*
Package model 用于模型层定义，所有db及cache对象封装均定义在这里。
只允许在这里添加对外暴露的接口
*/
package model

import (
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/model/feed"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/model/session"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/model/user"
)

func NewFeed() *feed.Feed {
	return &feed.Feed{}
}

func NewUser() *user.User {
	return &user.User{}
}

func NewSession() *session.Session {
	return &session.Session{}
}
