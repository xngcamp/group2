package skel

import (
	"encoding/json"
	"net/http"

	"camp/lib"
	"camp/Execample/api"
	"camp/Execample/service"

	"github.com/simplejia/clog/api"
)

// AddReq 定义输入
type AddReq struct {
	ID int64 `json:"id"`
}

// Regular 用于参数校验
func (addReq *AddReq) Regular() (ok bool) {
	if addReq == nil {
		return
	}

	if addReq.ID <= 0 {
		return
	}

	ok = true
	return
}

// AddResp 定义输出
type AddResp struct {
}

// Register just for demo
// @postfilter("Boss")
func (skel *Skel) Add(w http.ResponseWriter, r *http.Request) {
	fun := "user.Skel.Register"

	var addReq *AddReq
	if err := json.Unmarshal(skel.ReadBody(r), &addReq); err != nil || !addReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, addReq)
		skel.ReplyFail(w, lib.CodePara)
		return
	}

	skelApi := api.NewSkel()
	skelApi.ID = addReq.ID
	if err := service.NewSkel().Add(skelApi); err != nil {
		clog.Error("%s user.Register err: %v, req: %v", fun, err, addReq)
		skel.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &AddResp{}
	skel.ReplyOk(w, resp)

	// 进行一些异步处理的工作
	go lib.Updates(skelApi, lib.ADD, nil)

	return
}
