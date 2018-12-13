package feed

import (
	"encoding/json"
	"github.com/simplejia/clog/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/service"
	"github.com/xngcamp/group2/day8/gaoyuyue/lib"
	"net/http"
)

type UpdateReq struct {
	Txt string
}

type UpdateResp struct {

}

// @postfilter("Boss")
// @prefilter("Auth")
func (f *Feed) Update(w http.ResponseWriter, r *http.Request)  {
	fun := "controller.feed.Update"
	updateReq := &UpdateReq{}
	if err := json.Unmarshal(f.ReadBody(r), updateReq); err != nil {
		clog.Error("%s param err: %v, req: %v", fun, err, updateReq)
		f.ReplyFail(w, lib.CodePara)
		return
	}
	feed := api.NewFeed()
	if err := service.NewFeed().Update(feed); err != nil {
		f.ReplyFail(w, lib.CodePara)
		return
	}
	updateResp := &UpdateResp{}
	f.ReplyOk(w, updateResp)
}