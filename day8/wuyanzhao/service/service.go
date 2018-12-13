/*
Package service 用于定义服务层代码。
只允许在这里添加对外暴露的接口
*/
package service

import "camp/user/service/user"

func NewUserService() *user.UserService {
	return &user.UserService{}
}
