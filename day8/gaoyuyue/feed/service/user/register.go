package user

import (
	"errors"
	"fmt"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/model"
	"gopkg.in/mgo.v2/bson"
)

func (u *User) Register(userApi *api.User) (err error) {
	userModel := model.NewUser()
	userModel.Email = userApi.Email
	err = userModel.GetByEmail()
	if err == nil {
		err = errors.New("this email is registered")
		return
	}
	if err.Error() != "not found" {
		return
	}
	password, err := u.GetPassword(userApi)
	fmt.Println(password)
	if err != nil {
		return
	}
	userModel.Password = password
	userModel.Sex = userApi.Sex
	userModel.Nick = userApi.Nick
	userModel.Id = bson.NewObjectId()
	fmt.Println(userModel)
	if err = userModel.Add(); err != nil {
		return
	}
	return
}
