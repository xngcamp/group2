package feed

import (
	"camp/feed/api"
	"camp/feed/mongo"
	"github.com/globalsign/mgo"
)

type Feed api.Feed

//返回db
func (feed *Feed) Db() string {
	return "feed"
}

//返回table
func (feed Feed) Table() string{
	return "feed"
}

//返回db,col
func (feed Feed) GetC() (c *mgo.Collection) {
	db , talbe := feed.Db(),feed.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(talbe)
	return
}
