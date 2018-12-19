package redis

import (
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"strconv"
	"time"
	"videostore4/movie"
)

var (
	redisAddress string
	pool *redis.Pool
)

//初始化
func init()  {
	flag.StringVar(&redisAddress,"s","127.0.0.1:6379","redis server address")
	flag.Parse()

	poolSize := 20
	pool = &redis.Pool{
		MaxIdle: poolSize,
		IdleTimeout: time.Minute,
		Dial:func() (conn redis.Conn,err error) {
			conn, err = redis.Dial("tcp",redisAddress)
			if err != nil {
				return nil, err
			}
			return conn,nil
		},
	}
	CreateMovies()
}

func GetRedisConn() redis.Conn {
	return pool.Get()
}

func CreateMovies()  {
	conn := GetRedisConn()
	defer conn.Close()
	for _,movie := range movie.MovieList {
		movieTitle := movie.Title
		moviePriceCode := movie.PriceCode
		_,err := conn.Do("SET",movieTitle,strconv.Itoa(moviePriceCode))
		if err != nil {
			log.Fatal(err)
		}

	}
}

func GetMoviePriceCode(movieTitle string) (moviePriceCode int){
	conn := GetRedisConn()
	defer conn.Close()
	_priceCode,err := redis.String(conn.Do("GET",movieTitle))
	fmt.Println("_priceCode",_priceCode)
	if err != nil {
		log.Fatal(err)
	}
	moviePriceCode,err = strconv.Atoi(_priceCode)
	fmt.Println("moviePriceCode",moviePriceCode)
	if err != nil {
		log.Fatal(err)
	}
	return
}


