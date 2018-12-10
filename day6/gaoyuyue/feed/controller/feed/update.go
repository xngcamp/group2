package feed

import (
	"encoding/json"
	"github.com/simplejia/clog/api"
	"github.com/xngcamp/group2/day6/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day6/gaoyuyue/feed/service"
	"github.com/xngcamp/group2/day6/gaoyuyue/lib"
	"net/http"
)

// @postfilter("Cors")
func (f *Feed) Update(w http.ResponseWriter, r *http.Request)  {
	fun := "controller.feed.Update"
	var feed api.Feed
	if err := json.Unmarshal(f.ReadBody(r), &feed); err != nil {
		clog.Error("%s param err: %v, req: %v", fun, err, feed)
		f.ReplyFail(w, lib.CodePara)
		return
	}
	if err := service.NewFeed().Update(&feed); err != nil {
		f.ReplyFail(w, lib.CodePara)
		return
	}
}