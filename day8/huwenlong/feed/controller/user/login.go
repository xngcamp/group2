package user

import (
	"camp/feed/service"
	"camp/lib"
	"encoding/json"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/simplejia/clog/api"
	"net/http"
)

type LoginReq struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (loginReq *LoginReq) Regular() (ok bool) {
	if loginReq == nil || loginReq.Email == "" || loginReq.Password == ""{
		return
	}
	//验证邮箱格式
	//if...
	ok = true
	return
}

type LoginResp struct {
	Token bson.ObjectId `json:"token"`
}

//@postfilter("Boss")
func (user *User) Login(w http.ResponseWriter, r *http.Request)  {
	fun := "user.User.Login"

	var loginReq *LoginReq
	fmt.Println("pre",loginReq)
	if err := json.Unmarshal(user.ReadBody(r),&loginReq); err != nil {
		clog.Error("%s param err: %v, req: %v", fun, err, loginReq)
		user.ReplyFail(w,lib.CodeSrv)
		return
	}

	fmt.Println("post",loginReq)
	userApi,err := service.NewUser().Login(loginReq.Email,loginReq.Password)

	if err != nil {
		clog.Error("%s param err: %v, req: %v", fun, err, loginReq)
		user.ReplyFail(w,lib.CodeSrv)
		return
	}

	loginResp := &LoginResp{
		userApi.UserId,
	}

	user.ReplyOk(w,loginResp)

	go lib.Updates(userApi,lib.GET,nil)

	return
}
