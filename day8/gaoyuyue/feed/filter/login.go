package filter

import (
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/service"
	"github.com/xngcamp/group2/day8/gaoyuyue/lib"
	"net/http"
)

//type TokenBody struct {
//	Token string `json:"token"`
//	Body interface{} `json:"body"`
//}

//todo 将token放置在http header便于Login filter处理
func Login(w http.ResponseWriter, r *http.Request, p map[string]interface{}) bool {
	c := p["__C__"].(lib.IBase)
	//tokenBody := &TokenBody{}
	//if err := json.Unmarshal(c.ReadBody(r), tokenBody); err != nil {
	//	return false
	//}
	sessionService := service.NewSession()
	sessionApi, err := sessionService.GetById(r.Header.Get("token"))
	if err != nil {
		return false
	}
	c.SetParam("session", sessionApi)
	//c.SetParam(lib.KeyBody, tokenBody.Body)
	return true
}
