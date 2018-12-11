package feed

import (
	"camp/feed/service"
	"net/http"
)
// Add just for demo
// @postfilter("Boss")
func (feed *Feed) FeedList(w http.ResponseWriter, r *http.Request)  {
	fed :=service.NewFeed()
	w.Header().Set("Content-Type", "applcation/json; charset=UTF-8")
	w.Write(fed.FeedList())

}