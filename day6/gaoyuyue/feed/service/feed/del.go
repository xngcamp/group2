package feed

import (
	"github.com/xngcamp/group2/day6/gaoyuyue/feed/model"
	"gopkg.in/mgo.v2/bson"
)

func (f *Feed) DelById(id string) error {
	modelFeed := model.NewFeed()
	modelFeed.Id = bson.ObjectIdHex(id)
	return modelFeed.Del()
}
