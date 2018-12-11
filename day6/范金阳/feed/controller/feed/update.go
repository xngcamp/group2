package feed

import (
	"camp/feed/api"
	"camp/feed/service"
	"camp/lib"
	"encoding/json"
	"fmt"
	"github.com/simplejia/clog/api"
	"net/http"
	"strings"
)



// Regular 用于参数校验
func (addReq *AddReq) Regular_update() (ok bool) {
	if addReq == nil {
		return
	}

	//验证为空
	if addReq.Txt == ""{
		return
	}
	//验证空格
	if strings.Trim(addReq.Txt," ") == "" {
		return
	}

	ok = true
	return
}



// Add just for demo
// @postfilter("Boss")
func (feed *Feed) UpdateFeed(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.update"

	var addReq *AddReq
	if err := json.Unmarshal(feed.ReadBody(r), &addReq); err != nil || !addReq.Regular_update() {
		clog.Error("%s param err: %v, req: %v", fun, err, addReq)
		feed.ReplyFail(w, lib.CodePara)
		fmt.Print(err)
		return
	}

	feedApi := api.NewFeed()
	feedApi.ID=addReq.Id
	feedApi.TXT=addReq.Txt
	if err := service.NewFeed().UpdateFeed(feedApi); err != nil {
		clog.Error("%s skel.Add err: %v, req: %v", fun, err, addReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &AddResp{}
	feed.ReplyOk(w, resp)

	// 进行一些异步处理的工作
	go lib.Updates(feedApi, lib.ADD, nil)

	return
}
