package user

import (
	"camp/feed/api"
	"camp/feed/model"
	"errors"
	"fmt"
)

//err返回值有问题
func (user *User) Login(email string, password string) (userApi *api.User,err error) {
	userModel := model.NewUser()
	userModel.Email = email
	//先检测数据库中是否存在
	if _,_err := userModel.GetEmail();_err != nil {
		return
	}
	userModel,err = userModel.Login()
	fmt.Println("userPassword",userModel.Password,"----password",password)
	if err != nil {
		return
	}
	if userModel.Password != password {
		err = errors.New("password is wrong")
		return
	}
	if userModel == nil {
		return
	}
	userApi = (*api.User)(userModel)
	return
}
