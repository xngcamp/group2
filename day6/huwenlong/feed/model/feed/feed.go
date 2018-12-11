package feed

import (
	"camp/feed/api"
	"camp/feed/mongo"
	"github.com/globalsign/mgo"
)

type Feed api.Feed

//Db 返回 db 对应的类型
func (feed *Feed) Db() (db string) {
	return "feed"
}

//Table 返回table name
func (feed *Feed) Table() (table string) {
	return "feed"
}
//GetC 返回db name
func (feed *Feed) GetC() (c *mgo.Collection) {
	db,table := feed.Db(), feed.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}
