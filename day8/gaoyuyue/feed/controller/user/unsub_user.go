package user

import (
	"encoding/json"
	"github.com/simplejia/clog/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/service"
	"github.com/xngcamp/group2/day8/gaoyuyue/lib"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type UnsubReq struct {
	NeedUnsubUid bson.ObjectId `json:"need_unsub_uid"`
}

type UnsubResp struct {

}

// @prefilter("Auth")
// @postfilter("Cors","Boss")
func (u *User) UnsubUser(w http.ResponseWriter, r *http.Request) {
	fn := "controller.user.UnsubUser"
	unsubReq := &UnsubReq{}
	if err := json.Unmarshal(u.ReadBody(r), unsubReq); err != nil {
		u.ReplyFail(w, lib.CodePara)
		clog.Error(fn + "-> Error: %v", err)
		return
	}
	//session, ok := u.GetParam("session")
	token, ok := u.GetParam("token")
	if !ok {
		u.ReplyFail(w, lib.CodePara)
		clog.Error(fn + "-> Error: don't get session")
		return
	}
	//userId := session.(*api.Session).UserId
	userId := token.(*api.Token).UserId
	userService := service.NewUser()
	if err := userService.DelSub(userId, unsubReq.NeedUnsubUid); err != nil {
		u.ReplyFail(w, lib.CodeSrv)
		clog.Error(fn + "-> Error: %v", err)
		return
	}
	unsubResp := &UnsubResp{}
	u.ReplyOk(w, unsubResp)
}
