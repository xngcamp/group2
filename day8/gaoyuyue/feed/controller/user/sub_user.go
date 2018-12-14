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

type SubReq struct {
	NeedSubUid bson.ObjectId `json:"need_sub_uid"`
}

type SubResp struct {

}

// @prefilter("Auth")
// @postfilter("Cors","Boss")
func (u *User) SubUser(w http.ResponseWriter, r *http.Request) {
	fn := "controller.user.SubUser"
	subReq := &SubReq{}
	if err := json.Unmarshal(u.ReadBody(r), subReq); err != nil {
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
	subId := subReq.NeedSubUid
	if userId == subId {
		u.ReplyFail(w, lib.CodePara)
		u.ReplyFailWithDetail(w, lib.CodePara, "不能关注自己")
		clog.Error(fn + "-> Error: userId == subId")
		return
	}
 	if err := userService.AddSub(userId, subId); err != nil {
		u.ReplyFail(w, lib.CodeSrv)
		clog.Error(fn + "-> Error: %v", err)
		return
	}
	subResp := &SubResp{}
	u.ReplyOk(w, subResp)
}