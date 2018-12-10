package feed

import (
	"camp/feed/api"
	"camp/feed/model"
)

// Get 定义获取操作
func (feed *Feed) Get(id string) (feedApi *api.Feed, err error) {
	feedModel := model.NewFeed()
	feedModel.ID = id
	if feedModel, err = feedModel.Get(); err != nil {
		return
	}

	if feedModel == nil {

		return
	}

	feedApi = &api.Feed{
		ID:  feedModel.ID,
		Txt: feedModel.Txt,
	}

	return
}
