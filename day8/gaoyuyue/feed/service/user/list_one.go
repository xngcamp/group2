package user

import (
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/model"
)

func (u *User) ListOne(userApi *api.User) (err error) {
	userModel := model.NewUser()
	userModel.Id = userApi.Id
	err = userModel.GetById()
	if err != nil {
		return
	}
	userApi.Nick = userModel.Nick
	userApi.Email = userModel.Email
	userApi.Sex = userModel.Sex
	return
}