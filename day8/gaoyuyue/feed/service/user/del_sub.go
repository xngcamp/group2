package user

import (
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/model"
	"gopkg.in/mgo.v2/bson"
)

func (u *User) DelSub(userId bson.ObjectId, subUserId bson.ObjectId) (err error) {
	subUserModel := model.NewUser()
	subUserModel.Id = subUserId
	if err = subUserModel.GetById(); err != nil {
		return
	}
	userModel := model.NewUser()
	userModel.Id = userId
	if err = userModel.DelSub(subUserId); err != nil {
		return
	}
	if err = subUserModel.DelFans(userId); err != nil {
		return
	}
	return
}

