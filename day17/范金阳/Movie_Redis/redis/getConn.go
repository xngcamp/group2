package redis

import "github.com/garyburd/redigo/redis"

func GetConn() (c redis.Conn){
	c, err := redis.Dial("tcp", "127.0.0.1:6380")
	if err != nil {
		c = nil
		return
	}
	return
}
