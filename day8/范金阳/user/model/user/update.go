package user

import "github.com/globalsign/mgo/bson"

func (user *User) Update() (err error) {
	c := user.GetC()
	defer c.Database.Session.Close()
	if err = c.Update(bson.M{"_id":user.Token},user);err !=nil {
		return
	}
	return
}
