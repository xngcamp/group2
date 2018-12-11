package feed

import (
	"camp/feed/api"
	"camp/feed/util"
	"encoding/json"
)

// 定义查询操作
func (feed *Feed) FeedList() []byte {
	utiFeed := util.NewFeed()
	list := utiFeed.FeedList()
	db :=make([]api.Feed,0)
	//遍历数据放入数组
	for _,f :=range list{
		db = append(db,f)
	}
	//fmt.Println("db===>>",db)
	data , _ :=json.Marshal(db)


	return data
}
