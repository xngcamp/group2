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

type UpdateResp struct {
}

// @postfilter("Cors")
func (u *User) Update(w http.ResponseWriter, r *http.Request)  {
	fn := "controller.user.Update"
	updateReq := &UpdateReq{UpdateInfo:&api.User{}}
	if err := json.Unmarshal(u.ReadBody(r), updateReq); err != nil {
		u.ReplyFail(w, lib.CodePara)
		fmt.Println(err)
		clog.Error(fn + "-> Error: %v", err)
		return
	}
	fmt.Println(updateReq.UpdateInfo)
	sessionService := service.NewSession()
	sessionApi, err := sessionService.GetById(updateReq.Token)
	if err != nil {
		u.ReplyFail(w, lib.CodePara)
		fmt.Println(err)
		clog.Error(fn+"-> Error: %v", err)
		return
	}
	updateReq.UpdateInfo.Id = sessionApi.UserId
	fmt.Println(updateReq.UpdateInfo)
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