package filter

import (
	"fmt"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/service"
	"github.com/xngcamp/group2/day8/gaoyuyue/lib"
	"net/http"
)

//todo 将token放置在http header便于Login filter处理
func Auth(w http.ResponseWriter, r *http.Request, p map[string]interface{}) bool {
	fmt.Println("Auth......")
	c := p["__C__"].(lib.IBase)
	sessionService := service.NewSession()
	cookie, e := r.Cookie("token")
	if e != nil {
		fmt.Println(e)
		return false
	}
	sessionApi, err := sessionService.GetById(cookie.Value)
	fmt.Println(sessionApi)
	if err != nil {
		fmt.Println(err)
		return false
	}
	c.SetParam("session", sessionApi)
	return true
}
