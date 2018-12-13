package user

import (
	"github.com/globalsign/mgo/bson"
)

// Login 定义用户登录操作
func (userModel *UserModel) Login(userInfo *UserModel) (err error) {
	c := userModel.GetC()
	defer c.Database.Session.Close()

	//根据用户名查询出用户信息保存在userInfo中
	err = c.Find(bson.M{"_email": userModel.Email}).One(userInfo)

	if err != nil {
		return
	}

	return
}

//根据用户名密码查询出用户信息，保存到loginUser中
func (loginModel *LoginModel) InsertLogin() (err error) {
	c := loginModel.GetC()
	defer c.Database.Session.Close()

	err = c.Insert(loginModel)

	if err != nil {
		return
	}

	return
}
