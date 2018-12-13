package user

import (
	"encoding/json"
	"fmt"
	"github.com/simplejia/clog/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/service"
	"github.com/xngcamp/group2/day8/gaoyuyue/lib"
	"net/http"
)

type UpdateReq struct {
	Token string `json:"token"`
	UpdateInfo *api.User `json:"update_info"`
}

type UpdateInfo api.User

type UpdateResp struct {
}

// @postfilter("Cors","Boss")
// @prefilter("Auth")
func (u *User) Update(w http.ResponseWriter, r *http.Request)  {
	fn := "controller.user.Update"
	updateReq := &UpdateReq{UpdateInfo:&api.User{}}
	if err := json.Unmarshal(u.ReadBody(r), updateReq); err != nil {
		u.ReplyFail(w, lib.CodePara)
		fmt.Println(err)
		clog.Error(fn + "-> Error: %v", err)
		return
	}
	session, ok := u.GetParam("session")
	if !ok {
		u.ReplyFail(w, lib.CodePara)
		clog.Error(fn+"-> Error: don't get session")
		return
	}

	updateReq.UpdateInfo.Id = session.(*api.Session).UserId
	userService := service.NewUser()
	if err := userService.Update(updateReq.UpdateInfo); err != nil {
		fmt.Println(err)
		u.ReplyFail(w, lib.CodePara)
		clog.Error(fn+"-> Error: %v", err)
		return
	}
	updateResp := &UpdateResp{}
	u.ReplyOk(w, updateResp)
}
