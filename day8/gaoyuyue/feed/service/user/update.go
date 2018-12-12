package user

import (
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/model"
)

func (u *User) Update(userApi *api.User) (err error) {
	userModel := model.NewUser()
	userModel.Id = userApi.Id
	err = userModel.GetById()
	if err != nil {
		return
	}
	if userApi.Sex != 0 {
		userModel.Sex = userApi.Sex
	}
	if userApi.Nick != "" {
		userModel.Nick = userApi.Nick
	}
	userApi.Email = userModel.Email
	if userApi.Password != "" {
		password, err := u.GetPassword(userApi)
		if err != nil {
			return err
		}
		userModel.Password = password
	}
	err = userModel.Update()
	if err != nil {
		return
	}
	return
}
