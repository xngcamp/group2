package user

import (
	"camp/user/api"
	"camp/user/model"
)

func (user *User) FindOne(userApi *api.User) (apiUser *api.User, err error)  {
	userModel := model.NewUser()
	userModel.Email = userApi.Email
	userModel.Password = userApi.Password
	userModel,err = userModel.FindOne()
	if err != nil {
		return
	}
	if userModel == nil {
		return
	}

	if userModel.Password != userApi.Password {
		return
	}
	apiUser = (*api.User)(userModel)
	return
}
