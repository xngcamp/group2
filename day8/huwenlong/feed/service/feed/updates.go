package feed

import (
	"camp/feed/api"
	"camp/feed/model"
)

//Add定义新增操作
func (feed *Feed) Update(feedApi *api.Feed) (err error)  {
	feedModel := model.NewFeed()
	feedModel.Id = feedApi.Id
	_,err = feedModel.Get()

	feedModel.Txt = feedApi.Txt
	if err = feedModel.Update(); err != nil{
		return
	}
	return
}
