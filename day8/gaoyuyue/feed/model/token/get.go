package token

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2/bson"
)

func (t *Token) Get() (err error) {
	conn := t.GetConn()
	v, e := redis.String(conn.Do("GET", t.Id.Hex()))
	if e != nil {
		return e
	}
	t.UserId = bson.ObjectIdHex(v)
	return
}
