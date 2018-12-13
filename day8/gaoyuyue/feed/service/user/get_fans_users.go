package user

import (
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/model"
	"gopkg.in/mgo.v2/bson"
)

func (u *User) GetFansUsers(userId bson.ObjectId) (apiUsers []*api.User, err error) {
	userModel := model.NewUser()
	userModel.Id = userId
	err = userModel.GetById()
	if err != nil {
		return
	}
	apiUsers = make([]*api.User, 0)
	modelUsers, e := userModel.GetUsersByIds(userModel.Followers)
	if e != nil {
		err = e
		return
	}
	for _, u := range modelUsers {
		userApi := api.NewUser()
		userApi.Id = u.Id
		userApi.Email = u.Email
		userApi.Nick = u.Nick
		userApi.Sex = u.Sex
		apiUsers = append(apiUsers, userApi)
	}
	return
}
