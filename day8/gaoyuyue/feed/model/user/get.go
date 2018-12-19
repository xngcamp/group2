package user

import "gopkg.in/mgo.v2/bson"

func (u *User) GetByEmail() (err error) {
	c := u.GetC()
	defer c.Database.Session.Close()
	if err = c.Find(bson.M{"email": u.Email}).One(u); err != nil {
		return
	}
	return
}

func (u *User) GetById() (err error) {
	c := u.GetC()
	defer c.Database.Session.Close()
	if err = c.Find(bson.M{"_id": u.Id}).One(u); err != nil {
		return
	}
	return
}