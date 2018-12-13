package user

import (
	"camp/feed/mongo"
	"camp/user/api"
	"github.com/globalsign/mgo"
)

type User api.User

func (user *User) Db() string{
	return "user"
}

func (user *User) Table() string{
	return "user"
}

func (user *User) GetC() (c *mgo.Collection){
	db , table := user.Db(),user.Table()

	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)

	return c
}
