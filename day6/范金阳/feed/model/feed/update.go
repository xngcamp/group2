package feed

import (
	"camp/feed/api"
	"camp/feed/util"
)

// 定义更新操作
func (feed *Feed) UpdateFeed(fed *api.Feed) (err error) {
	utiFeed := util.NewFeed()
	utiFeed.UpdateFeed(fed)



	return
}
