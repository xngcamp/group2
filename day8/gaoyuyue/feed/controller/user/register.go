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

type RegisterReq struct {
	Nick string
	Sex int
	Email string
	Password string
}

func (registerReq *RegisterReq) Regular() (ok bool) {
	//todo emial正则匹配，nick字数限制
	if registerReq == nil || registerReq.Email == "" || registerReq.Password == "" ||
		registerReq.Sex < 1 || registerReq.Sex > 2 || registerReq.Nick == ""{
		return
	}
	return true
}

type RegisterResp struct {

}

// @postfilter("Cors")
func (u *User) Register(w http.ResponseWriter, r *http.Request)  {
	fn := "controller.user.Register"

	addReq := &RegisterReq{}
	if err := json.Unmarshal(u.ReadBody(r), addReq); err != nil || !addReq.Regular() {
		u.ReplyFail(w, lib.CodePara)
		clog.Error(fn + "-> Error: %v", err)
		return
	}

	userApi := api.NewUser()
	userApi.Email = addReq.Email
	userApi.Nick = addReq.Nick
	userApi.Password = addReq.Password
	userApi.Sex = addReq.Sex

	fmt.Println(userApi)
	userService := service.NewUser()
	if err := userService.Register(userApi); err != nil {
		u.ReplyFailWithDetail(w, lib.CodePara, err.Error())
		clog.Error(fn + "-> Error: %v", err)
		return
	}

	resp := &RegisterResp{}
	u.ReplyOk(w, resp)
}
