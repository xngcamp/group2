/*
Package service 用于定义服务层代码。
只允许在这里添加对外暴露的接口
*/
package service

import (
	"github.com/xngcamp/group2/day6/gaoyuyue/feed/service/feed"
)

func NewFeed() *feed.Feed {
	return &feed.Feed{}
}