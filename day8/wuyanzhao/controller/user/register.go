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

//注册请求类
type RegisterReq struct {
	Nick     string `json:"nick" bson:"_nick"`         //用户名
	Sex      byte   `json:"sex" bson:"_sex"`           //性别 1|2
	Email    string `json:"email" bson:"_email"`       //电子邮箱
	Password string `json:"password" bson:"_password"` //密码
}

// Regular 用于参数校验
func (registerReq *RegisterReq) Regular() (ok bool) {
	if registerReq == nil {
		return
	}

	if len(strings.TrimSpace(registerReq.Nick)) == 0 {
		return
	}

	if registerReq.Sex != 1 && registerReq.Sex != 2 {
		return
	}

	//这里需要判断email格式，先粗略这么写
	if len(strings.TrimSpace(registerReq.Email)) == 0 {
		return
	}

	if len(strings.TrimSpace(registerReq.Password)) == 0 {
		return
	}

	ok = true
	return
}

// RegisterResp 定义输出
type RegisterResp struct {
}

// Add just for demo
// @postfilter("Boss")
func (user *User) Register(w http.ResponseWriter, r *http.Request) {
	fun := "user.UserController.Register"

	var registerReq *RegisterReq
	if err := json.Unmarshal(user.ReadBody(r), &registerReq); err != nil || !registerReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, registerReq)
		user.ReplyFail(w, lib.CodePara)
		return
	}

	userApi := api.NewUserApi()
	userApi.Nick = registerReq.Nick
	userApi.Sex = registerReq.Sex
	userApi.Email = registerReq.Email
	userApi.Password = registerReq.Password
	if err := service.NewUserService().Register(userApi); err != nil {
		//log.Println("controller register err:", err)
		clog.Error("%s user.Register err: %v, req: %v", fun, err, registerReq)
		//user.ReplyFail(w, lib.CodeSrv)
		user.ReplyFailWithDetail(w, lib.CodeSrv, err.Error())
		return
	}

	resp := &RegisterResp{}
	user.ReplyOk(w, resp)

	// 进行一些异步处理的工作
	go lib.Updates(userApi, lib.ADD, nil)

	return
}
