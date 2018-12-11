package feed

import (
	"camp/feed/api"
	"camp/feed/model"
	"github.com/globalsign/mgo/bson"
)

//Add定义新增操作
func (feed *Feed) Add(feedApi *api.Feed) (err error)  {
	feedModel := model.NewFeed()
	feedModel.Txt = feedApi.Txt
	feedModel.Id = bson.NewObjectId()
	if err = feedModel.Add(); err != nil{
		return
	}
	return
}
