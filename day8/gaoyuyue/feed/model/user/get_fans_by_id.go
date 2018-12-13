package user

import (
	"gopkg.in/mgo.v2/bson"
)

func (u *User) GetFansById() (fans []*User, err error) {
	c := u.GetC()
	defer c.Database.Session.Close()
	fans = make([]*User, 0)
	err = c.Find(bson.M{"following.0": u.Id}).All(&fans)
	if err != nil {
		return
	}
	return
}
