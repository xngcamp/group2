package user

import (
	"camp/user/api"
	"camp/user/model"
	"camp/user/model/user"
	"fmt"

	uuid "go.uuid-master"
)

func (userService *UserService) Login(loginApi *api.LoginApi) (usertoken string, err error) {
	//将登录用户信息查询出来，保存到userModel中
	userModel := model.NewUserModel()
	userModel.Email = loginApi.Email

	var userInfo user.UserModel //保存登录用户信息
	if err = userModel.Login(&userInfo); err != nil {
		return
	}

	//生成usertoken
	usertoken = uuid.Must(uuid.NewV4()).String()
	loginInfo := model.NewLoginUserModel()
	loginInfo.UserID = userInfo.ID
	loginInfo.Usertoken = usertoken
	fmt.Println(userInfo.ID)

	//将用户id和usertoken插入login 表
	if err = loginInfo.InsertLogin(); err != nil {
		return
	}

	return
}
