package feed

import (
	"camp/feed/model"
	"github.com/globalsign/mgo/bson"
)

//Del定义删除操作

func (feed *Feed) Del(id bson.ObjectId) (err error) {
	feedModel := model.NewFeed()
	feedModel.Id = id

	if err = feedModel.Del(); err != nil {
		return
	}
	return
}