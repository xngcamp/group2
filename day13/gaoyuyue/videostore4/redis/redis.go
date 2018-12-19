package redis

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	redisAddress = "127.0.0.1:6379"
	pool         *redis.Pool
)

func init() {
	poolSize := 20
	pool = &redis.Pool{
		MaxIdle:     poolSize,
		IdleTimeout: time.Minute,
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", redisAddress)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}

func GetRedisConn() redis.Conn {
	return pool.Get()
}