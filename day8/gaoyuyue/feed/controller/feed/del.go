package feed

import (
	"encoding/json"
	"github.com/simplejia/clog/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/service"
	"github.com/xngcamp/group2/day8/gaoyuyue/lib"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type DelReq struct {
	Id bson.ObjectId `json:"id"`
}

type DelResp struct {

}

// @postfilter("Boss")
// @prefilter("Auth")
func (f *Feed) Del(w http.ResponseWriter, r *http.Request)  {
	fn := "controller.feed.Del"
	delReq := &DelReq{}
	if err := json.Unmarshal(f.ReadBody(r), delReq); err != nil {
		clog.Error("%s param err: %v, req: %v", fn, err, delReq)
		f.ReplyFail(w, lib.CodePara)
		return
	}
	if err := service.NewFeed().DelById(delReq.Id); err != nil {
		f.ReplyFail(w, lib.CodePara)
		return
	}
	delResp := &DelResp{}
	f.ReplyOk(w, delResp)
}
