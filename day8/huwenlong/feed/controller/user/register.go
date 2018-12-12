package user

import (
	"camp/feed/api"
	"camp/feed/service"
	"camp/lib"
	"encoding/json"
	"github.com/simplejia/clog/api"
	"net/http"
)

type RegisterReq struct {
	Nick string `json:"nick"`
	Sex int `json:"sex"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (registerReq * RegisterReq) Ragular() (ok bool) {
	if registerReq == nil {
		return
	}
	ok = true
	return
}

type RegisterResp struct {
	Nick string `json:"nick"`
	Email string `json:"email"`
}

//@postfilter("Boss")
func (user *User) Register(w http.ResponseWriter, r *http.Request)  {
	fun := "user.User.Register"

	var registerReq * RegisterReq
	if err := json.Unmarshal(user.ReadBody(r), &registerReq); err != nil {
		clog.Error("json:--%s user.Register err: %v, req: %v", fun, err, registerReq)
		user.ReplyFail(w,lib.CodeSrv)
		return
	}

	userApi := api.NewUser()
	userApi.Nick = registerReq.Nick
	userApi.Sex = registerReq.Sex
	userApi.Email = registerReq.Email
	userApi.Password = registerReq.Password

	if err := service.NewUser().Register(userApi);err != nil {
		clog.Error("service:--%s user.Register err: %v, req: %v", fun, err, registerReq)
		user.ReplyFail(w,lib.CodeSrv)
		return
	}
	registerResp := &RegisterResp{
		userApi.Nick,
		userApi.Email,
	}
	user.ReplyOk(w,registerResp)

	go lib.Updates(userApi,lib.ADD,nil)
}