package user

import (
	"gopkg.in/mgo.v2/bson"
)

func (u *User) AddFans(fansId bson.ObjectId) (err error) {
	c := u.GetC()
	defer c.Database.Session.Close()
	err = c.UpdateId(u.Id, bson.M{"$push": bson.M{"followers": fansId}})
	if err != nil {
		return
	}
	return
}
