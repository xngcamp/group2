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

type ListOneReq struct {
	Uid bson.ObjectId `json:"uid"`
}

type ListOneResp struct {
	Id string `json:"id"`
	Nick string `json:"nick"`
	Email string `json:"email"`
	Sex int `json:"sex"`
}

// @prefilter("Auth")
// @postfilter("Boss","Cors")
func (u *User) ListOne(w http.ResponseWriter, r *http.Request)  {
	fn := "controller.user.ListOne"
	fmt.Println(fn)
	listOneReq := &ListOneReq{}
	if err := json.Unmarshal(u.ReadBody(r), listOneReq); err != nil {
		u.ReplyFail(w, lib.CodePara)
		clog.Error(fn + "-> Error: %v", err)
		return
	}
	fmt.Println(listOneReq)
	userService := service.NewUser()
	userApi := api.NewUser()
	userApi.Id = listOneReq.Uid
	fmt.Println("start list one")
	if err := userService.ListOne(userApi); err != nil {
		u.ReplyFail(w, lib.CodeSrv)
		clog.Error(fn + "-> Error: %v", err)
		return
	}
	fmt.Println("end list one")
	listOneResp := &ListOneResp{}
	listOneResp.Id = userApi.Id.Hex()
	listOneResp.Sex = userApi.Sex
	listOneResp.Nick = userApi.Nick
	listOneResp.Email = userApi.Email
	u.ReplyOk(w, listOneResp)
}