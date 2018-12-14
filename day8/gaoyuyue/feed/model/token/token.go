package token

import (
	"fmt"
	redis2 "github.com/garyburd/redigo/redis"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/redis"
)

type Token api.Token

func (t *Token) Business() string {
	return "token"
}

func (t *Token) GetConn() redis2.Conn {
	business := t.Business()
	pool := redis.RDS[business]
	fmt.Println(pool)
	return pool.Get()
}