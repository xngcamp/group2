package feed

import (
	"camp/feed/service"
	"camp/lib"
	"encoding/json"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/simplejia/clog/api"
	"net/http"
)

type DelReq struct {
	ID bson.ObjectId `json:"id"`
}
//接受的数据的判断
func (delReq *DelReq) Regular() (ok bool) {
	if delReq == nil{
		return
	}
	ok = true
	return
}

type DelResp struct {

}

//Del just for demo
//@postfilter("Boss")
func (feed *Feed) Del(w http.ResponseWriter,r *http.Request) {
	fun := "feed.Feed.Del"

	var delReq *DelReq
	if err := json.Unmarshal(feed.ReadBody(r), &delReq); err != nil || !delReq.Regular(){
		clog.Error("%s param err: %v, req: %v", fun, err, delReq)
		feed.ReplyFail(w,lib.CodeSrv)
		return
	}

	fmt.Println("get_id",delReq.ID)
	feedApi, err := service.NewFeed().Get(delReq.ID)
	if err != nil{
		clog.Error("%s param err: %v, req: %v", fun, err, delReq)
		feed.ReplyFail(w,lib.CodeSrv)
		return
	}

	//if feedApi.Id == nil {
	//	deltail := "feed not found"
	//	feed.ReplyFailWithDetail(w,lib.CodePara,deltail)
	//	return
	//}

	if err := service.NewFeed().Del(delReq.ID); err != nil{
		clog.Error("%s param err: %v, req: %v", fun, err, delReq)
		feed.ReplyFail(w,lib.CodeSrv)
		return
	}

	resp := &DelResp{}
	feed.ReplyOk(w,resp)

	//进行一些异步处理的操作
	go lib.Updates(feedApi,lib.DELETE,nil)
	return
}