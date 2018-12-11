package feed

import (
	"camp/feed/api"
	"camp/feed/model"
)

func (feed *Feed) UpdateFeed(feedApi *api.Feed) (err error) {
	feedModel := model.NewFeed()
	if err = feedModel.UpdateFeed(feedApi); err != nil {
		return
	}

	return
}