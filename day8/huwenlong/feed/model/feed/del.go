package feed

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func (feed *Feed) Del() (err error)  {
	c := feed.GetC()
	defer c.Database.Session.Close()
	//根据id查找feed 进行删除操作
	//bson.M
	err = c.Remove(bson.M{"_id":feed.Id})
	if err != nil {
		if err != mgo.ErrNotFound {
			return
		}
		err = nil
		return
	}
	return
}