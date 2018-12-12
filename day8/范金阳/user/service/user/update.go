package user

import (
	"camp/user/api"
	"camp/user/model"
)

func (user *User) Update(apiUser *api.User) (err error) {
	userModel := model.NewUser()
	userModel.Token = apiUser.Token
	userModel,err = userModel.FindOne()
	if err != nil || userModel==nil{
		return
	}
	modelUser := model.NewUser()
	modelUser.Token = userModel.Token
	modelUser.Email = userModel.Email
	if apiUser.Nick == "" {
		modelUser.Nick = userModel.Nick
	}else {
		modelUser.Nick = apiUser.Nick
	}
	if apiUser.Password == "" {
		modelUser.Password = userModel.Password
	}else {
		modelUser.Password = apiUser.Password
	}
	if apiUser.Sex == 0 {
		modelUser.Sex = userModel.Sex
	}else {
		modelUser.Sex = apiUser.Sex
	}
	err = modelUser.Update()
	if err != nil {
		return
	}

	return
}
