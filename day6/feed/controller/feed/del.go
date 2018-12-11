package feed

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"camp/feed/service"
	"camp/lib"

	"github.com/simplejia/clog/api"
)

// DelReq 定义输入
type DelReq struct {
	ID  string `json:"id"`
	Txt string `json:"txt"`
}

// Regular 用于参数校验
func (delReq *DelReq) Regular() (ok bool) {
	if delReq == nil {
		return
	}

	tempID, _ := strconv.Atoi(delReq.ID)
	if tempID <= 0 {
		return
	}

	ok = true
	return
}

// DelResp 定义输出
type DelResp struct {
}

// Del just for demo
// @postfilter("Boss")
func (feed *Feed) Del(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.Del"

	var delReq *DelReq
	if err := json.Unmarshal(feed.ReadBody(r), &delReq); err != nil || !delReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, delReq)
		feed.ReplyFail(w, lib.CodePara)
		return
	}

	feedApi, err := service.NewFeed().Get(strings.TrimSpace(delReq.ID))
	if err != nil {
		clog.Error("%s skel.Get err: %v, req: %v", fun, err, delReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	if feedApi == nil {
		detail := "feed not found"
		feed.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}

	if err := service.NewFeed().Del(delReq.ID, delReq.Txt); err != nil {
		clog.Error("%s skel.Del err: %v, req: %v", fun, err, delReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &DelResp{}
	feed.ReplyOk(w, resp)

	// 进行一些异步处理的工作
	go lib.Updates(feedApi, lib.DELETE, nil)

	return
}
