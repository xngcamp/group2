package feed

import (
	"github.com/xngcamp/group2/day6/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day6/gaoyuyue/feed/model"
)

func (f *Feed) Update(feed *api.Feed) error {
	modelFeed := model.NewFeed()
	if err := modelFeed.GetById(feed.Id); err != nil {
		return err
	}
	modelFeed.Content = feed.Content
	return modelFeed.Update()
}
