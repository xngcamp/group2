package user

import (
	"github.com/simplejia/clog/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/service"
	"github.com/xngcamp/group2/day8/gaoyuyue/lib"
	"net/http"
)

type SubUsersResp struct {
	Subs []*api.User `json:"subs"`
}

// @prefilter("Auth")
// @postfilter("Boss","Cors")
func (u *User) GetSubUsers(w http.ResponseWriter, r *http.Request)  {
	fn := "controller.user.GetSubUsers"
	session, ok := u.GetParam("session")
	if !ok {
		u.ReplyFail(w, lib.CodePara)
		clog.Error(fn + "-> Error: don't get session")
		return
	}
	userId := session.(*api.Session).UserId
	userService := service.NewUser()
	apiUsers, err := userService.GetSubUsers(userId)
	if err != nil {
		u.ReplyFail(w, lib.CodeSrv)
		clog.Error(fn + "-> Error: %v", err)
		return
	}
	subUsersResp := &SubUsersResp{}
	subUsersResp.Subs = apiUsers
	u.ReplyOk(w, subUsersResp)
}
