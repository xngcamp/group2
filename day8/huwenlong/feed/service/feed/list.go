package feed

import (
	"camp/feed/api"
	"camp/feed/model"
)

//getall 定义获取全部
func (feed *Feed) List() (feeds []*api.Feed,err error)  {
	modelFeeds, err := model.NewFeed().List()
	feeds = make([]*api.Feed,0,len(modelFeeds))
	if err != nil {
		return
	}
	for _,feed := range modelFeeds {
		feeds  = append(feeds,(*api.Feed)(feed))
	}
	return
}
