package feed

import (
	"github.com/xngcamp/group2/day6/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day6/gaoyuyue/feed/model"
)

func (f *Feed) List() (feeds []*api.Feed, err error) {
	modelFeeds, err := model.NewFeed().List()
	feeds = make([]*api.Feed, 0,len(modelFeeds))
	if err != nil {
		return
	}
	for _, feed := range modelFeeds {
		feeds = append(feeds, (*api.Feed)(feed))
	}
	return
}
