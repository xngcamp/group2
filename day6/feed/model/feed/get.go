package feed

import (
	"fmt"
	"strings"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Get 定义获取操作
func (feed *Feed) Get() (feedRet *Feed, err error) {
	feedRet = &Feed{}
	c := feed.GetC()
	defer c.Database.Session.Close()
	tempid := strings.TrimSpace(feed.ID)
	fmt.Printf("tempid: %v\n", tempid)
	err = c.Find(bson.M{"_id": tempid}).One(&feedRet)
	fmt.Printf("After find %v\n", feedRet)

	if err != nil {
		fmt.Errorf("%v\n", err.Error())
		if err != mgo.ErrNotFound {
			return
		}
		err = nil
		return
	}

	return
}
