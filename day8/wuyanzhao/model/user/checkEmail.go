package user

import (
	"errors"

	"github.com/globalsign/mgo/bson"
)

// Register 定义用户注册操作
func (userModel *UserModel) CheckEmail(email string) (err error, count int) {
	c := userModel.GetC()
	defer c.Database.Session.Close()

	count, err = c.Find(bson.M{"_email": email}).Count()
	if err != nil {
		return
	} else if count > 0 {
		return errors.New("The email has existed,please try another."), count
	}

	return
}
