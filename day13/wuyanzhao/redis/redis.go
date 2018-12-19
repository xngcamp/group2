package myredis

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var Pool *redis.Pool

func init() {
	redisAddress := "127.0.0.1:6379"
	maxIdle := 20
	Pool = &redis.Pool{
		MaxIdle:     maxIdle,
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
