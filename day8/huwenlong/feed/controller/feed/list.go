package feed

import (
	"camp/feed/service"
	"camp/lib"
	"net/http"
)

func (feed *Feed) List(w http.ResponseWriter,r *http.Request)  {
	feeds, err := service.NewFeed().List()
	if err != nil {
		feed.ReplyFail(w,lib.CodeSrv)
		return
	}
	feed.ReplyOk(w,feeds)
}



