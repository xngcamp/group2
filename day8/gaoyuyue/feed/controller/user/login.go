package user

import (
	"encoding/json"
	"github.com/simplejia/clog/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/service"
	"github.com/xngcamp/group2/day8/gaoyuyue/lib"
	"net/http"
)

type LoginReq struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (loginReq *LoginReq) Regular() (ok bool) {
	//todo
	if loginReq == nil || loginReq.Email == "" || loginReq.Password == "" {
		return
	}
	return true
}

type LoginResp struct {
	Token string `json:"token"`
}

// @postfilter("Boss")
func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	fn := "controller.user.Login"
	loginReq := &LoginReq{}
	if err := json.Unmarshal(u.ReadBody(r),loginReq); err != nil || !loginReq.Regular() {
		u.ReplyFail(w, lib.CodePara)
		clog.Error(fn + "-> Error: %v", err)
		return
	}
	userApi := api.NewUser()
	userApi.Email = loginReq.Email
	userApi.Password = loginReq.Password
	userService := service.NewUser()
	userId, err := userService.Login(userApi)
	if err != nil {
		u.ReplyFailWithDetail(w, lib.CodePara, err.Error())
		clog.Error(fn + "-> Error: %v", err)
		return
	}
	sessionService := service.NewSession()
	sessionApi, err := sessionService.Add(userId)
	if err != nil {
		u.ReplyFail(w, lib.CodePara)
		clog.Error(fn + "-> Error: %v", err)
		return
	}
	token := sessionApi.Id.Hex()
	loginResp := &LoginResp{}
	loginResp.Token = token
	u.ReplyOk(w, loginResp)
}
