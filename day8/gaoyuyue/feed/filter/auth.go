package filter

import (
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/service"
	"github.com/xngcamp/group2/day8/gaoyuyue/lib"
	"net/http"
)

//todo 将token放置在http header便于Login filter处理
func Auth(w http.ResponseWriter, r *http.Request, p map[string]interface{}) bool {
	c := p["__C__"].(lib.IBase)
	sessionService := service.NewSession()
	cookie, e := r.Cookie("token")
	if e != nil {
		return false
	}
	sessionApi, err := sessionService.GetById(cookie.Value)
	if err != nil {
		return false
	}
	c.SetParam("session", sessionApi)
	return true
}
