/*
Package model 用于模型层定义，所有db及cache对象封装均定义在这里。
只允许在这里添加对外暴露的接口
*/
package model

import "camp/user/model/user"

// NewUserModel 构造UserModel对象
func NewUserModel() *user.UserModel {
	return &user.UserModel{}
}

// NewLoginUserModel 构造LoginUserModel对象
func NewLoginUserModel() *user.LoginModel {
	return &user.LoginModel{}
}
