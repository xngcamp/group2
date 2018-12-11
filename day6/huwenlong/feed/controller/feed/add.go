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

type AddReq struct {
	Txt string `json:"txt"`
	Id bson.ObjectId `json:"id"`
}
// Regular 用于参数校验
func (addReq *AddReq) Regular() (ok bool) {
	if addReq == nil {
		return
	}
	
	ok = true
	return
}
//响应中传递的结构体
type AddResp struct {
	Txt string `json:"ret_txt"`
}

//Add just for demo
//@postfilter("Boss")
func (feed *Feed) Add( w http.ResponseWriter,r *http.Request)  {
	fun := "feed.Feed.Add"
	
	var addReq *AddReq
	if err := json.Unmarshal(feed.ReadBody(r),&addReq);err != nil || !addReq.Regular(){
		clog.Error("%s feed.Add err: %v, req: %v", fun, err, addReq)
		feed.ReplyFail(w,lib.CodeSrv)
		return
	}
	
	feedApi := api.NewFeed()
	feedApi.Txt = addReq.Txt

	if err := service.NewFeed().Add(feedApi); err != nil {
		clog.Error("%s feed.Add err: %v, req: %v", fun, err, addReq)
		feed.ReplyFail(w,lib.CodeSrv)
		return
	}
	resp := &AddResp{
		feedApi.Txt,
	}
	feed.ReplyOk(w,resp)

	//进行一些异步处理的操作
	go lib.Updates(feedApi,lib.ADD,nil)

	return

}
