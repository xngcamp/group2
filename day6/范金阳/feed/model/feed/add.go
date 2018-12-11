package feed

import (
	"camp/feed/api"
	"camp/feed/util"
)

// Add 定义新增操作
func (feed *Feed) Add(fed *api.Feed) (err error) {
	utiFeed := util.NewFeed()
	id := utiFeed.GeneraeId()
	utiFeed.AddMap(id,fed)



	return
}

// Add 定义新增操作
//func (feed *Feed) Add(fed *api.Feed) (err error) {
//	c := feed.GetC()
//	defer c.Database.Session.Close()
//	err = c.Insert(feed)
//	if err!=nil {
//		return
//	}
//
//
//
//	return
//}

