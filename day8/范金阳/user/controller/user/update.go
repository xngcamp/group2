package user

import (
	"camp/lib"
	"camp/user/api"
	"camp/user/service"
	"encoding/json"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/simplejia/clog/api"
	"net/http"
	"strings"
)

type UpdateReq struct {
	Nick string `json:"nick"`	//用户姓名
	Sex int `json:"sex"`	//用户性别（1|2）
	Email string `json:"email"` 	//用户邮箱，做登录用，并且检测是否已经注册
	Password string `json:"password"` 	//用户的密码（后期加密）
	Token bson.ObjectId `json:"token" bson:"_id"`
}

//数据校验方法
func (addReq *UpdateReq) Regular() (ok bool){
	if addReq == nil {
		fmt.Println(">>>>>>>>>>>1")
		return
	}
	
	ok = true
	return
}

type UpdateResp struct {

}


func (user *User) Update(w http.ResponseWriter, r *http.Request) {
	fun := "user.Controller.add"
	var updateReq *UpdateReq
	if err :=json.Unmarshal(user.ReadBody(r),&updateReq);err != nil || !updateReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, updateReq)
		user.ReplyFail(w,lib.CodePara)
		return
	}
	apiUser := api.NewUser()
	apiUser.Email = updateReq.Email
	apiUser.Sex = updateReq.Sex
	apiUser.Nick = updateReq.Nick
	apiUser.Password = updateReq.Password
	apiUser.Token = updateReq.Token

	if err := service.NewUser().Update(apiUser); err != nil {
		clog.Error("%s user.Register err: %v, req: %v", fun, err,updateReq)
		user.ReplyFail(w , lib.CodeSrv)
		return
	}
	addResp := &AddResp{}
	user.ReplyOk(w,addResp)

	go lib.Updates(apiUser, lib.ADD, nil)

	return
}

