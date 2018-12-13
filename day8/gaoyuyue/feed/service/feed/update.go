package feed

import (
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/model"
)

func (f *Feed) Update(feed *api.Feed) (err error) {
	modelFeed := model.NewFeed()
	if err = modelFeed.GetById(feed.Id); err != nil {
		return
	}
	modelFeed.Content = feed.Content
	err = modelFeed.Update()
	if err != nil {
		return
	}
	return
}
