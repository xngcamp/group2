package user

import (
	"fmt"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/model"
	"gopkg.in/mgo.v2/bson"
)

func (u *User) AddSub(userId bson.ObjectId, subUserId bson.ObjectId) (err error) {
	subUserModel := model.NewUser()
	subUserModel.Id = subUserId
	if err = subUserModel.GetById(); err != nil {
		return
	}
	userModel := model.NewUser()
	fmt.Println(userId)
	userModel.Id = userId
	fmt.Println(userModel)
	if err = userModel.AddSub(subUserId); err != nil {
		return
	}
	if err = subUserModel.AddFans(userId); err != nil {
		return
	}
	return
}
