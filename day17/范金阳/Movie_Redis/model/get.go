package model

import (
	"Movie_Redis/api"
	redis3 "Movie_Redis/redis"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

func(movie *Movie) Get() (mov *api.Movie) {
	c := redis3.GetConn()
	defer c.Close()
	if c == nil{
		return
	}
	data ,err :=redis.String(c.Do("get",movie.Id))
	if err != nil {
		return
	}
	b :=[]byte(data)
	err =json.Unmarshal(b,&mov)
	if err != nil {
		return
	}

	return
}
