package user

import (
	"camp/feed/api"
	"camp/feed/model"
	"github.com/globalsign/mgo/bson"
)

func (user *User) Register(userApi *api.User) (err error) {
	userModel := model.NewUser()
	userModel.UserId = bson.NewObjectId()
	userModel.Nick = userApi.Nick
	userModel.Sex = userApi.Sex
	userModel.Email = userApi.Email
	userModel.Password = userApi.Password

	//检测邮箱是否以注册
	if _,err = userModel.GetEmail();err == nil {
		return
	}
	if err = userModel.Register(); err != nil {
		return
	}
	return
}


