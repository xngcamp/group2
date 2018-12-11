package feed

import (
	"camp/feed/api"
	"camp/feed/model"
)

// Add 定义新增操作
func (feed *Feed) Add(feedApi *api.Feed) (err error) {
	//u_feed := util.NewFeed()
	feedModel := model.NewFeed()
	feedModel.TXT=feedApi.TXT
	feedModel.ID=feedApi.ID
	//feedModel.Id=u_feed.GeneraeId()
	if err = feedModel.Add(feedApi); err != nil {
		return
	}

	return
}