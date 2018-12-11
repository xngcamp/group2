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
	//初始化数据库与表的名称
	db,table := feed.Db(), feed.Table()
	//连接数据库
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	//创建表
	c = sessionCopy.DB(db).C(table)
	//返回当前表的集合
	return
}
