package feed

import (
	"camp/feed/api"
	"camp/feed/service"
	"camp/lib"
	"encoding/json"
	"fmt"
	"github.com/simplejia/clog/api"
	"net/http"
)

func (feed Feed) Del(w  http.ResponseWriter,r *http.Request)  {
	fun :="feed.controller.del"
	var delReq *AddReq
	err := json.Unmarshal(feed.ReadBody(r),&delReq)
	if err != nil{
		clog.Error("%s param err: %v, req: %v", fun, err, delReq)
		feed.ReplyFail(w, lib.CodePara)
		fmt.Print(err)
		return
	}

	feedApi := api.NewFeed()
	feedApi.ID = delReq.Id
	service.NewFeed().Del(feedApi)
}
