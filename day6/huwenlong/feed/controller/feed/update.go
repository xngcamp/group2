package feed

import (
	"camp/feed/service"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"net/http"

	"camp/lib"
	"camp/feed/api"

	"github.com/simplejia/clog/api"

)

type UpdateReq struct {
	Id bson.ObjectId `json:"id"`
	Txt string `json:"txt"`
}
// Regular 用于参数校验
func (updateReq *UpdateReq) Regular() (ok bool) {
	if updateReq == nil {
		return
	}

	ok = true
	return
}
type UpdateResp struct {
	Id bson.ObjectId `json:"ret_id"`
	Txt string `json:"ret_txt"`
}

//Add just for demo
//@postfilter("Boss")
func (feed *Feed) Update( w http.ResponseWriter,r *http.Request)  {
	fun := "feed.Feed.Add"

	var updateReq *UpdateReq
	if err := json.Unmarshal(feed.ReadBody(r),&updateReq);err != nil || !updateReq.Regular(){
		clog.Error("%s feed.Add err: %v, req: %v", fun, err, updateReq)
		feed.ReplyFail(w,lib.CodeSrv)
		return
	}

	feedApi := api.NewFeed()
	feedApi.Txt = updateReq.Txt
	feedApi.Id = updateReq.Id
	if err := service.NewFeed().Update(feedApi); err != nil {
		clog.Error("%s feed.Add err: %v, req: %v", fun, err, updateReq)
		feed.ReplyFail(w,lib.CodeSrv)
		return
	}
	resp := &AddResp{
		feedApi.Id,
		feedApi.Txt,
	}
	feed.ReplyOk(w,resp)

	//进行一些异步处理的操作
	go lib.Updates(feedApi,lib.ADD,nil)

	return

}
