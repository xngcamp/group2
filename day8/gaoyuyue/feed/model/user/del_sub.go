package user

import "gopkg.in/mgo.v2/bson"

func (u *User) DelSub(subId bson.ObjectId) (err error)  {
	c := u.GetC()
	defer c.Database.Session.Close()
	err = c.UpdateId(u.Id, bson.M{"$pull": bson.M{ "following": subId}})
	if err != nil {
		return
	}
	return
}
