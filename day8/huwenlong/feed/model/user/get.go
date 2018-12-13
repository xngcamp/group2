package user

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func (user *User) GetId(userid bson.ObjectId) (userApi *User,err error) {
	c := user.GetC()

	defer c.Database.Session.Close()

	err = c.Find(bson.M{"_id": userid}).One(&userApi)
	if err != nil {
		if err != mgo.ErrNotFound {
			fmt.Println("not found")
			return
		}
		return
	}
	return
}
func (user *User) GetEmail() (userApi *User,err error) {
	c := user.GetC()

	defer c.Database.Session.Close()

	err = c.Find(bson.M{"_id":user.Email}).One(&userApi)
	if err != nil {
		if err != mgo.ErrNotFound {
			fmt.Println("not found")
			return
		}
		return
	}
	return
}