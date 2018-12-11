package feed

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//Get定义获取操作
func (feed *Feed) Get() (feedRet *Feed,err error)  {
	c := feed.GetC()
	defer c.Database.Session.Close()

	err = c.Find(bson.M{"_id":feed.Id}).One(&feedRet)

	if err != nil{
		if err != mgo.ErrNotFound{
			return
		}
		err = nil
		return
	}
	return
}
