package user

import (
	"gopkg.in/mgo.v2/bson"
)

func (u *User) GetUsersByIds(ids []bson.ObjectId) (users []*User, err error) {
	c := u.GetC()
	defer c.Database.Session.Close()
	users = make([]*User, 0)
	err = c.Find(bson.M{"_id": bson.M{"$in": ids}}).All(&users)
	if err != nil {
		return
	}
	return
}
