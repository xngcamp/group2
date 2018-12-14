package user

import (
	"github.com/simplejia/clog/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/service"
	"github.com/xngcamp/group2/day8/gaoyuyue/lib"
	"net/http"
)

type FansUsersResp struct {
	Fans []*api.User `json:"fans"`
}

// @prefilter("Auth")
// @postfilter("Boss","Cors")
func (u *User) GetFansUsers(w http.ResponseWriter, r *http.Request)  {
	fn := "controller.user.GetFanUsers"

	//session, ok := u.GetParam("session")
	token, ok := u.GetParam("token")
	if !ok {
		u.ReplyFail(w, lib.CodePara)
		clog.Error(fn + "-> Error: don't get session")
		return
	}
	//userId := session.(*api.Session).UserId
	userId := token.(*api.Token).UserId
	userService := service.NewUser()
	apiUsers, err := userService.GetFansUsers(userId)
	if err != nil {
		u.ReplyFail(w, lib.CodeSrv)
		clog.Error(fn + "-> Error: %v", err)
		return
	}
	fansUsersResp := &FansUsersResp{}
	fansUsersResp.Fans = apiUsers
	u.ReplyOk(w, fansUsersResp)
}