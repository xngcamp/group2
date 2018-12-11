package feed

import (
	"camp/feed/api"
	"camp/feed/model"
)

func (feed *Feed) Del(apiFeed *api.Feed){
	feedModel :=model.NewFeed()
	feedModel.Del(apiFeed)
}
