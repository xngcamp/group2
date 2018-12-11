package feed

import (
	"encoding/json"
	"net/http"
	"strconv"

	"camp/feed/api"
	"camp/feed/service"
	"camp/lib"

	"github.com/simplejia/clog/api"
)

// AddReq 定义输入
type AddReq struct {
	ID  string `json:"id"`
	Txt string `json:"txt"`
}

// Regular 用于参数校验
func (addReq *AddReq) Regular() (ok bool) {
	if addReq == nil {
		return
	}

	tempID, _ := strconv.Atoi(addReq.ID)
	if tempID <= 0 {
		return
	}

	// if len(strings.Trim(addReq.Txt, " ")) == 0 {
	// 	return
	// }

	ok = true
	return
}

// AddResp 定义输出
type AddResp struct {
}

// Add just for demo
// @postfilter("Boss")
func (feed *Feed) Add(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.Add"

	var addReq *AddReq
	if err := json.Unmarshal(feed.ReadBody(r), &addReq); err != nil || !addReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, addReq)
		feed.ReplyFail(w, lib.CodePara)
		return
	}

	feedApi := api.NewFeed()
	feedApi.ID = addReq.ID
	feedApi.Txt = addReq.Txt
	if err := service.NewFeed().Add(feedApi); err != nil {
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
