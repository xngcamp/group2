package feed

import (
	"camp/feed/api"
	"camp/feed/util"
)

// Add 定义新增操作
func (feed *Feed) Del(fed *api.Feed) (err error) {
	utiFeed := util.NewFeed()
	utiFeed.Del(fed)



	return
}


