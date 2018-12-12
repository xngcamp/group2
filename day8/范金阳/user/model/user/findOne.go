package user

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func (user *User) FindOne()(userSet *User, err error){
	c := user.GetC()


	defer c.Database.Session.Close()
	err = c.Find(bson.M{"email":user.Email}).One(&userSet)
	if err != nil{
		if err != mgo.ErrNotFound {
			return
		}
		err = nil
		return
	}
	return
}
