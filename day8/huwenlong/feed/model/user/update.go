package user

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
)

func (user *User) Update(userid bson.ObjectId) (err error) {
	c := user.GetC()

	defer c.Database.Session.Close()
	q := bson.M{}
	q["_password"] = user.Password
	q["_sex"] = user.Sex
	q["_nick"] = user.Nick
	q["_email"] = user.Email
	fmt.Println("q------",q)
	err = c.UpdateId(userid,q)

	if err != nil {
		return
	}
	return
}
