package user

import (
	"camp/lib"
	"camp/user/api"
	"camp/user/service"
	"encoding/json"
	"net/http"
	"strings"

	clog "github.com/simplejia/clog/api"
)

//登录请求类
type LoginReq struct {
	Email    string `json:"email" bson:"_email"`       //电子邮箱
	Password string `json:"password" bson:"_password"` //密码
}

// Regular 用于参数校验
func (loginReq *LoginReq) Regular() (ok bool) {
	if loginReq == nil {
		return
	}

	//这里需要判断email格式，先粗略这么写
	if len(strings.TrimSpace(loginReq.Email)) == 0 {
		return
	}

	if len(strings.TrimSpace(loginReq.Password)) == 0 {
		return
	}

	ok = true
	return
}

// LoginResp 定义输出
type LoginResp struct {
	Usertoken string `json:"usertoken" bson:"_usertoken"`
}

// Add just for demo
// @postfilter("Boss")
func (user *User) Login(w http.ResponseWriter, r *http.Request) {
	fun := "user.UserController.Login"

	var loginReq *LoginReq
	if err := json.Unmarshal(user.ReadBody(r), &loginReq); err != nil || !loginReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, loginReq)
		user.ReplyFailWithDetail(w, lib.CodePara, err.Error())
		return
	}

	loginService := service.NewUserService()
	loginApi := api.LoginApi{}
	loginApi.Email = loginReq.Email
	loginApi.Password = loginReq.Password
	usertoken, err := loginService.Login(&loginApi)
	if err != nil {
		clog.Error("%s user.Login err: %v, req: %v", fun, err, loginReq)
		user.ReplyFailWithDetail(w, lib.CodeSrv, err.Error())
		return
	}

	resp := &LoginResp{usertoken}
	user.ReplyOk(w, resp)

	// 进行一些异步处理的工作
	go lib.Updates(loginApi, lib.ADD, nil)

	return
}
