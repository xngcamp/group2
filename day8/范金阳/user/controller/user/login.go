package user

import (
	"camp/lib"
	"camp/user/api"
	"camp/user/service"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/simplejia/clog/api"
	"net/http"
	"strings"
)
//输入参数
type LoginReq struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (loginReq *LoginReq) Regular() (ok bool) {
	if loginReq == nil {
		return
	}

	if loginReq.Email == "" || strings.Trim(loginReq.Email," ") == "" {
		return
	}

	if loginReq.Password == "" || strings.Trim(loginReq.Password," ") == "" {
		return
	}

	ok = true
	return
}

type LoginResp struct {
	//apiUser *api.User `json:"user，omitempty"`
	//Nick string `json:"nick"`	//用户姓名
	//Sex int `json:"sex"`	//用户性别（1|2）
	//Email string `json:"email"` 	//用户邮箱，做登录用，并且检测是否已经注册
	//Password string `json:"password"` 	//用户的密码（后期加密）
	Token bson.ObjectId `json:"token"`
}

func (user *User) Login(w http.ResponseWriter, r *http.Request){
	fun := "user.Controller.login"
	var loginReq *LoginReq
	err := json.Unmarshal(user.ReadBody(r),&loginReq)
	if err != nil || !loginReq.Regular(){
		clog.Error("%s param err: %v, req: %v", fun, err, loginReq)
		user.ReplyFail(w, lib.CodePara)
		return
	}
	apiUser := api.NewUser()
	apiUser.Password = loginReq.Password
	apiUser.Email = loginReq.Email
	userApi , err := service.NewUser().FindOne(apiUser)
	if err != nil {
		clog.Error("%s param err: %v, req: %v", fun, err, loginReq)
		user.ReplyFail(w, lib.CodeSrv)
		return
	}
	if userApi==nil{
		clog.Error("%s param err: %v, req: %v", fun, err, loginReq)
		user.ReplyFail(w, lib.CodeDataNo)
		return
	}
	var resp = &LoginResp{
		//apiUser:userApi,
		//userApi.Nick,
		//userApi.Sex,
		//userApi.Email,
		//userApi.Password,
		userApi.Token,
	}
	user.ReplyOk(w,resp)
	// 进行一些异步处理的工作
	go lib.Updates(userApi, lib.GET, nil)

	return
}