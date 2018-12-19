package feed

import (
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/model"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func (f *Feed) Add(feed *api.Feed) error {
	modelFeed := model.NewFeed()
	feed.Id = bson.NewObjectId()
	feed.CreateTime = time.Now()
	modelFeed.Id = feed.Id
	modelFeed.Content = feed.Content
	modelFeed.CreateTime = feed.CreateTime
	return modelFeed.Add()
}
