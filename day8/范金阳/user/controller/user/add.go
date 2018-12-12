package user

import (
	"camp/lib"
	"camp/user/api"
	"camp/user/service"
	"encoding/json"
	"fmt"
	"github.com/simplejia/clog/api"
	"net/http"
	"strings"
)

type AddReq struct {
	Nick string `json:"nick"`	//用户姓名
	Sex int `json:"sex"`	//用户性别（1|2）
	Email string `json:"email"` 	//用户邮箱，做登录用，并且检测是否已经注册
	Password string `json:"password"` 	//用户的密码（后期加密）
}

//数据校验方法
func (addReq *AddReq) Regular() (ok bool){
	if addReq == nil {
		fmt.Println(">>>>>>>>>>>1")
		return
	}
	if addReq.Sex != 1 && addReq.Sex!= 2 {
		fmt.Println(">>>>>>>>>>>2")
		return
	}
	if addReq.Password ==""||strings.Trim(addReq.Password," ")=="" {
		fmt.Println(">>>>>>>>>>>3")
		return
	}
	if addReq.Nick ==""||strings.Trim(addReq.Nick," ")=="" {
		fmt.Println(">>>>>>>>>>>4")
		return
	}
	if addReq.Email ==""||strings.Trim(addReq.Email," ")=="" {
		fmt.Println(">>>>>>>>>>>5")
		return
	}
	ok = true
	return
}

type AddResp struct {

}

func (user *User) Register(w http.ResponseWriter, r *http.Request) {
	fun := "user.Controller.add"
	var addReq *AddReq
	if err :=json.Unmarshal(user.ReadBody(r),&addReq);err != nil || !addReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, addReq)
		user.ReplyFail(w,lib.CodePara)
		return
	}
	apiUser := api.NewUser()
	apiUser.Email = addReq.Email
	apiUser.Sex = addReq.Sex
	apiUser.Nick = addReq.Nick
	apiUser.Password = addReq.Password

	if err := service.NewUser().Register(apiUser); err != nil {
		clog.Error("%s user.Register err: %v, req: %v", fun, err,addReq)
		user.ReplyFail(w , lib.CodeSrv)
		return
	}
	addResp := &AddResp{}
	user.ReplyOk(w,addResp)

	go lib.Updates(apiUser, lib.ADD, nil)

	return
}


