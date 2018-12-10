package feed

import (
	"fmt"
	"github.com/xngcamp/group2/day6/gaoyuyue/feed/service"
	"github.com/xngcamp/group2/day6/gaoyuyue/lib"
	"net/http"
)

// @postfilter("Cors")
func (f *Feed) Del(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	id := r.Form.Get("id")
	fmt.Println(id)
	if err := service.NewFeed().DelById(id); err != nil {
		fmt.Println(err)
		f.ReplyFail(w, lib.CodePara)
		return
	}
}
