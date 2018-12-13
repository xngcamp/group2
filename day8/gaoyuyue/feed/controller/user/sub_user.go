package user

import (
	"encoding/json"
	"fmt"
	"github.com/simplejia/clog/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/service"
	"github.com/xngcamp/group2/day8/gaoyuyue/lib"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type SubReq struct {
	NeedSubUid bson.ObjectId `json:"need_sub_uid"`
}

type SubResp struct {

}

// @prefilter("Auth")
// @postfilter("Boss")
func (u *User) SubUser(w http.ResponseWriter, r *http.Request) {
	fn := "controller.user.SubUser"
	fmt.Println(fn)
	subReq := &SubReq{}
	if err := json.Unmarshal(u.ReadBody(r), subReq); err != nil {
		u.ReplyFail(w, lib.CodePara)
		clog.Error(fn + "-> Error: %v", err)
		return
	}
	session, ok := u.GetParam("session")
	fmt.Println(session)
	if !ok {
		u.ReplyFail(w, lib.CodePara)
		clog.Error(fn + "-> Error: don't get session")
		return
	}
	fmt.Println("start add sub")
	userService := service.NewUser()
	fmt.Println("session-----")
	fmt.Println(session.(*api.Session))
	fmt.Println(session.(*api.Session).UserId)
	if err := userService.AddSub(session.(*api.Session).UserId, subReq.NeedSubUid); err != nil {
		fmt.Println(err)
		u.ReplyFail(w, lib.CodeSrv)
		clog.Error(fn + "-> Error: %v", err)
		return
	}
	fmt.Println("end add sub")
	subResp := &SubResp{}
	u.ReplyOk(w, subResp)
}