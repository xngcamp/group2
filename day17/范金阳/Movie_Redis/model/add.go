package model

import (
	"Movie_Redis/redis"
	"encoding/json"
)

func (movie *Movie) Add() (ok bool) {
	c := redis.GetConn()
	if c == nil {
		return
	}
	defer c.Close()
	date , _ := json.Marshal(movie)
	_,err := c.Do("set",movie.Id,date)
	if err != nil{
		return
	}
	ok = true
	return
}
