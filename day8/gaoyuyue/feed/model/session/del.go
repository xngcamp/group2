package session

import "gopkg.in/mgo.v2/bson"

func (s *Session) Del() (err error) {
	c := s.GetC()
	defer c.Database.Session.Close()

	err = c.Remove(bson.M{"_id": s.Id})
	if err != nil {
		return
	}
	return
}
