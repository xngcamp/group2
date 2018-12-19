package feed

import (
	"gopkg.in/mgo.v2/bson"
)

func (f *Feed) GetById(id bson.ObjectId) error {
	c := f.GetC()
	defer c.Database.Session.Close()
	return c.FindId(id).One(f)
}
