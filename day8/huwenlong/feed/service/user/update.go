package user

import (
	"camp/feed/api"
	"camp/feed/model"
	"fmt"
	"github.com/globalsign/mgo/bson"
)

func (user *User) Update(userApi *api.User,userid bson.ObjectId) (err error) {
	userModel := model.NewUser()

	fmt.Println("1111111111111")
	_user,_err:= userModel.GetId(userid)
	if _err != nil {
		fmt.Println("_err",_err)
		return
	}
	fmt.Println("2222222222222222222")
	userModel.Email = _user.Email
	//更新用户信息
	if userApi.Password == "" {
		userModel.Password = _user.Password
	}else {
		userModel.Password = userApi.Password
	}
	if userApi.Nick == "" {
		userModel.Nick = _user.Nick
	}else {
		userModel.Nick = userApi.Nick
	}
	if userApi.Sex == 0 {
		userModel.Sex = _user.Sex
	}else {
		userModel.Sex = userApi.Sex
	}
	fmt.Println("userModelInfo:",userModel)
	if err = userModel.Update(userid);err != nil {
		return
	}
	return
}