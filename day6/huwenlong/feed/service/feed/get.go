package feed

import (
	"camp/feed/api"
	"camp/feed/model"
	"fmt"
)

//Get定义获取操作
func (feed *Feed) Get(id int64) (feedApi *api.Feed,err error) {
	feedModel := model.NewFeed()
	feedModel.Id = id

	if feedModel,err = feedModel.Get(); err != nil{
		fmt.Println("feedModel is get error")
		return
	}
	if feedModel == nil{
		fmt.Println("feedModel is nil")
		return
	}
	feedApi = (*api.Feed)(feedModel)
	return
}