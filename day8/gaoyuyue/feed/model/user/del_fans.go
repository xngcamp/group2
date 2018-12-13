package user

import "gopkg.in/mgo.v2/bson"

func (u *User) DelFans(fansId bson.ObjectId) (err error) {
	c := u.GetC()
	defer c.Database.Session.Close()
	err = c.UpdateId(u.Id, bson.M{"$pull": bson.M{"followers": fansId}})
	if err != nil {
		return
	}
	return
}
