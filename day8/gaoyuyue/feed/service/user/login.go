package user

import (
	"errors"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/model"
	"gopkg.in/mgo.v2/bson"
)

func (u *User) Login(userApi *api.User) (userId bson.ObjectId, err error) {
	userModel := model.NewUser()
	userModel.Email = userApi.Email
	err = userModel.GetByEmail()
	if err != nil {
		err = errors.New("user is not exist")
		return
	}
	password, err := u.GetPassword(userApi)
	if err != nil {
		return
	}
	if userModel.Password != password {
		err = errors.New("password is invalid")
		return
	}
	userId = userModel.Id
	return
}