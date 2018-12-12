package feed

import (
	"github.com/globalsign/mgo"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/mongo"
)

type Feed api.Feed

func (f *Feed) Db() (db string) {
	return "feed"
}

func (f *Feed) Table() (table string) {
	return "feeds"
}

func (f *Feed) GetC() (c *mgo.Collection) {
	db, table := f.Db(), f.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}