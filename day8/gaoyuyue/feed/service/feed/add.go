package feed

import (
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/model"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func (f *Feed) Add(feedApi *api.Feed) (err error) {
	feedModel := model.NewFeed()
	feedApi.Id = bson.NewObjectId()
	feedApi.CreateTime = time.Now()
	feedModel.Id = feedApi.Id
	feedModel.Content = feedApi.Content
	feedModel.UserId = feedApi.UserId
	feedModel.CreateTime = feedApi.CreateTime
	err = feedModel.Add()
	if err != nil {
		return
	}
	return
}
