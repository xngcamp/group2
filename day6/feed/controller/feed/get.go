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

// GetReq 定义输入
type GetReq struct {
	ID  string `json:"id"`
	Txt string `json:"txt"`
}

// Regular 用于参数校验
func (getReq *GetReq) Regular() (ok bool) {
	if getReq == nil {
		return
	}

	tempID, _ := strconv.Atoi(getReq.ID)
	if tempID <= 0 {
		return
	}

	ok = true
	return
}

// GetResp 定义输出
type GetResp struct {
	Feed *api.Feed `json:"feed,omitempty"`
}

// Get just for demo
// @postfilter("Boss")
func (feed *Feed) Get(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.Get"

	var getReq *GetReq
	if err := json.Unmarshal(feed.ReadBody(r), &getReq); err != nil || !getReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, getReq)
		feed.ReplyFail(w, lib.CodePara)
		return
	}

	feedApi, err := service.NewFeed().Get(getReq.ID)
	if err != nil {
		clog.Error("%s feed.Get err: %v, req: %v", fun, err, getReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &GetResp{
		Feed: feedApi,
	}
	feed.ReplyOk(w, resp)

	// 进行一些异步处理的工作
	go lib.Updates(feedApi, lib.GET, nil)

	return
}
