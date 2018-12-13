package user

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/simplejia/clog/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/lib"
	"net/http"
)

type SubReq struct {
	NeedSubUid bson.ObjectId `json:"need_sub_uid"`
}

type SubResp struct {

}

func (u *User) SubUser(w http.ResponseWriter, r *http.Request) {
	fn := "controller.user.SubUser"
	subReq := &SubReq{}
	if err := json.Unmarshal(u.ReadBody(r), subReq); err != nil {
		u.ReplyFail(w, lib.CodePara)
		clog.Error(fn + "-> Error: %v", err)
		return
	}

	subResp := &SubResp{}
	u.ReplyOk(w, subResp)
}