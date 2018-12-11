package feed

import (
	"camp/feed/api"
	"fmt"
)

var feed_map =make(map[int]api.Feed,0)
//数据中的id
var id int =0
//自动生成id
func (feed *Feed) GeneraeId() int {
	id++
	return id
}
//模拟数据库添加
func (feed *Feed) AddMap(id int, fed *api.Feed){
	var s = api.NewFeed()
	s.ID = id
	s.TXT = fed.TXT
	feed_map[id] = *s
	fmt.Print("feedUtil===>>",feed_map)
}
//模拟数据库更新
func (feed *Feed) UpdateFeed(fed *api.Feed)  {
	var s = api.NewFeed()
	s.ID = fed.ID
	s.TXT = fed.TXT
	feed_map[s.ID]=*s
}
//模拟数据库查询
func (feed *Feed)FeedList() map[int]api.Feed {
	f := api.Feed{1,"5"}
	feed_map[1]=f
	return feed_map
}
//模拟数据库删除
func (feed *Feed)Del(fed *api.Feed){
	delete(feed_map,fed.ID)
}