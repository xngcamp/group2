package feed

import (
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/model"
	"gopkg.in/mgo.v2/bson"
)

func (f *Feed) DelById(id bson.ObjectId) (err error) {
	modelFeed := model.NewFeed()
	modelFeed.Id = id
	err = modelFeed.Del()
	if err != nil {
		return
	}
	return
}
