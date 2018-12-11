/*
Package skel just for demo
*/
package feed

import (
	"camp/feed/api"
	"camp/feed/mongo"

	mgo "github.com/globalsign/mgo"
)

// Skel 定义db对应的类型
type Feed api.Feed

// Db 返回db name
func (feed *Feed) Db() (db string) {
	return "feed"
}

// Table 返回table name
func (feed *Feed) Table() (table string) {
	return "feed"
}

// GetC 返回db col
func (feed *Feed) GetC() (c *mgo.Collection) {
	db, table := feed.Db(), feed.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}
