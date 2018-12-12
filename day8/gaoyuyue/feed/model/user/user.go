package user

import (
	"github.com/globalsign/mgo"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/mongo"
)

type User api.User

func (u *User) Db() (db string) {
	return "feed"
}

func (u *User) Table() (table string) {
	return "users"
}

func (u *User) GetC() (c *mgo.Collection) {
	db, table := u.Db(), u.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}

