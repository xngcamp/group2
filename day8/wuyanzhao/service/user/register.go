package user

import (
	"camp/user/api"
	"camp/user/model"
	"log"
)

func (userService *UserService) Register(userApi *api.UserApi) (err error) {
	userModel := model.NewUserModel()
	userModel.Nick = userApi.Nick
	userModel.Sex = userApi.Sex
	userModel.Email = userApi.Email
	userModel.Password = userApi.Password

	var count int
	if err, count = userModel.CheckEmail(userModel.Email); err != nil {
		return
	} else if count > 0 {
		log.Println("count:", count)
		return
	}

	if err = userModel.Register(); err != nil {
		return
	}
	return
}
