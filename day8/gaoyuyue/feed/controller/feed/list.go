package feed

import (
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/service"
	"github.com/xngcamp/group2/day8/gaoyuyue/lib"
	"net/http"
)

// @postfilter("Cors")
func (f *Feed) List(w http.ResponseWriter, r *http.Request) {
	feeds, err := service.NewFeed().List()
	if err != nil {
		f.ReplyFail(w, lib.CodePara)
		return
	}
	if err != nil {
		f.ReplyFail(w, lib.CodePara)
		return
	}
	f.ReplyOk(w, feeds)
}
