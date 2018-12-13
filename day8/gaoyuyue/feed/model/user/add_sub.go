package user

import (
	"gopkg.in/mgo.v2/bson"
)

func (u *User) AddSub(subId bson.ObjectId) (err error) {
	c := u.GetC()
	defer c.Database.Session.Close()
	err = c.UpdateId(u.Id, bson.M{"$push": bson.M{"following": subId}})
	if err != nil {
		return
	}
	return
}
