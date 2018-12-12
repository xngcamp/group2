package user

import (
	"camp/user/api"
	"camp/user/model"
	"errors"
)

func (user *User) Register(apiUser *api.User) (err error){
	userModel := model.NewUser()
	userModel.Email = apiUser.Email
	//判断邮箱是否已经注册
	if userModel ,err =userModel.FindOne();err != nil {

		return
	}

	if userModel != nil {
		err =errors.New("账号已注册！")
		return
	}

	modelUser := model.NewUser()
	modelUser.Email = apiUser.Email
	modelUser.Nick = apiUser.Nick
	modelUser.Password = apiUser.Password
	modelUser.Sex = apiUser.Sex

	err = modelUser.Add()
	if err != nil {
		return
	}

	return
}
