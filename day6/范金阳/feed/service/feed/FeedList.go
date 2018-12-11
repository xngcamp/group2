package feed

import "camp/feed/model"

func (feed *Feed) FeedList() []byte {
	fed :=model.NewFeed()
	return fed.FeedList()
}
