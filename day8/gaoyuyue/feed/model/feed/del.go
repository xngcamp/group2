package feed

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
)

func (f *Feed) Del() error {
	c := f.GetC()
	defer c.Database.Session.Close()
	fmt.Println(f.Id)
	return c.Remove(bson.M{"_id" : f.Id})
}
