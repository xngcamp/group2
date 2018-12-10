package feed

import (
	"fmt"

	"github.com/globalsign/mgo"
)

// Del 定义删除操作
func (feed *Feed) Del() (err error) {
	c := feed.GetC()
	fmt.Printf("链接：%v\n", c)
	defer c.Database.Session.Close()
	fmt.Printf("del feed: %v\n", feed)
	err = c.Remove(feed)
	if err != nil {
		fmt.Errorf("%v\n", err)
		if err != mgo.ErrNotFound {
			return
		}
		err = nil
		return
	}

	return
}
