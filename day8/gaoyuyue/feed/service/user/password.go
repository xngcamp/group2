package user

import (
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/util"
)

// 密码密文 = sha1(email + password)
func (u *User) GetPassword(apiUser *api.User) (password string, err error) {
	password, err = util.Sha1Encryption(apiUser.Email + apiUser.Password)
	if err != nil {
		return
	}
	return
}
