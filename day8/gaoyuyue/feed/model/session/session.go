package session

import (
	"github.com/globalsign/mgo"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/mongo"
)

type Session api.Session

func (s *Session) Db() (db string) {
	return "feed"
}

func (s *Session) Table() (table string) {
	return "sessions"
}

func (s *Session) GetC() (c *mgo.Collection) {
	db, table := s.Db(), s.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}
