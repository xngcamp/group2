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

type GetReq struct {
	ID bson.ObjectId `json:"id"`
} 

func (getReq *GetReq) Regular() (ok bool) {
	if getReq == nil {
		return 
	}

	ok = true
	return 
}

type GetResp struct {
	//Feed *api.Feed `json:"feed,omitempty"`
	Id bson.ObjectId `json:"ret_id"`
	Txt string `json:"ret_txt"`
} 

//Get just for demo
//@postfilter("Boss")
func (feed *Feed) Get(w http.ResponseWriter,r *http.Request) {
	fun := "feed.Feed.Get"

	var getReq *GetReq
	if err := json.Unmarshal(feed.ReadBody(r),&getReq); err != nil || !getReq.Regular(){
		clog.Error("%s param err: %v, req: %v", fun, err, getReq)
		feed.ReplyFail(w,lib.CodeSrv)
		return
	}
	fmt.Println("1",getReq.ID)
	feedApi,err := service.NewFeed().Get(getReq.ID)
	fmt.Println("2",feedApi)
	if err != nil {
		clog.Error("%s param err: %v, req: %v", fun, err, getReq)
		feed.ReplyFail(w,lib.CodeSrv)
		return
	}

	resp := &GetResp{
		feedApi.Id,
		feedApi.Txt,
	}
	feed.ReplyOk(w,resp)

	go lib.Updates(feedApi,lib.GET, nil)

	return
}