/*
Package service 用于定义服务层代码。
只允许在这里添加对外暴露的接口
*/
package service

import  (
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/service/feed"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/service/session"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/service/user"
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