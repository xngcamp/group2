package feed

import (
	"encoding/json"
	"github.com/simplejia/clog/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/service"
	"github.com/xngcamp/group2/day8/gaoyuyue/lib"
	"net/http"
)

type AddReq struct {
	Txt string `json:"txt"`
}

// @postfilter("Cors","Boss")
// @prefilter("Auth")
func (f *Feed) Add(w http.ResponseWriter, r *http.Request)  {
	fn := "controller.feed.Add"
	addReq := &AddReq{}
	if err := json.Unmarshal(f.ReadBody(r), addReq); err != nil {
		clog.Error("%s param err: %v, req: %v", fn, err, addReq)
		f.ReplyFail(w, lib.CodePara)
		return
	}
	session, ok := f.GetParam("session")
	if !ok {
		f.ReplyFail(w, lib.CodePara)
		clog.Error(fn+"-> Error: don't get session")
		return
	}

	feed := api.NewFeed()
	feed.UserId = session.(*api.Session).UserId
	feed.Content = addReq.Txt
	if err := service.NewFeed().Add(feed); err != nil {
		f.ReplyFail(w, lib.CodeSrv)
		return
	}
	f.ReplyOk(w, feed)
}