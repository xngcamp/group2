package feed

import (
	"camp/feed/model"
)

// Del 定义删除操作
func (feed *Feed) Del(id string, txt string) (err error) {
	feedModel := model.NewFeed()
	feedModel.ID = id
	feedModel.Txt = txt
	if err = feedModel.Del(); err != nil {
		return
	}

	return
}
