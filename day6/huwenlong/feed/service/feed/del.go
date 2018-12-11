package feed

import "camp/feed/model"

//Del定义删除操作

func (feed *Feed) Del(id int64) (err error) {
	feedModel := model.NewFeed()
	feedModel.Id = id

	if err = feedModel.Del(); err != nil {
		return
	}
	return
}