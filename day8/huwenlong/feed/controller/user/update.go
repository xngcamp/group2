package user

import (
	"camp/feed/api"
	"camp/feed/service"
	"camp/lib"
	"encoding/json"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/simplejia/clog/api"
	"net/http"
)

type UpdateReq struct {
	UserId bson.ObjectId `json:"token"`
	AUser api.User`json:"a_user"`
	//Email string `json:"email"`
	//Nick string `json:"nick"`
	//Sex int `json:"sex"`
	//Password string `json:"password"`
}

func (updateReq *UpdateReq) Regular() (ok bool) {
	if updateReq == nil || updateReq.UserId ==""{
		fmt.Println("the token is nil")
		return
	}

	ok = true
	return
}

type UpdateResp struct {
	Status string `json:"ret_status"`
}

//@postfilter("Boss")
func (user *User) Update(w http.ResponseWriter, r *http.Request) {
	fun := "user.User.Update"

	updateReq  := &UpdateReq{}
	if err := json.Unmarshal(user.ReadBody(r),&updateReq);err != nil || !updateReq.Regular() {
		clog.Error("json:--%s user.Update err: %v, req: %v", fun, err, updateReq)
		user.ReplyFail(w,lib.CodeSrv)
		return
	}

	
	userApi := api.NewUser()
	userApi.Email = updateReq.AUser.Email
	userApi.Nick = updateReq.AUser.Nick
	userApi.Sex = updateReq.AUser.Sex
	userApi.Password = updateReq.AUser.Password
	userid := updateReq.UserId
	fmt.Println("userApi",userApi,"userId",userid)
	if err := service.NewUser().Update(userApi,userid);err != nil {
		clog.Error("service:--%s user.Update err: %v, req: %v", fun, err, updateReq)
		user.ReplyFail(w,lib.CodeSrv)
		return
	}

	updateResp := &UpdateResp{
		"ok",
	}

	user.ReplyOk(w,updateResp)

	go lib.Updates(user,lib.UPDATE,nil)

	return
}