package main

import (
	"camp/week2/day2/wuyanzhao/customer"
	"camp/week2/day2/wuyanzhao/movie"
	"camp/week2/day2/wuyanzhao/rental"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

func init() {
	var aCustomer customer.Customer
	conn := aCustomer.GetRedisConn()
	defer conn.Close()
	db := 2 //电影库
	conn.Do("select", db)
	for i := 0; i < 10; i++ {
		star := rand.Float64() * 10
		movie1 := movie.BaseMovie{
			Id:       i,
			Title:    "Movie_" + strconv.Itoa(i),
			Year:     time.Now().Unix(),
			Author:   "wuyanzhao",
			Star:     star,
			TypeCode: movie.NORMAL,
		}
		if b, err := json.Marshal(movie1); err == nil {
			fmt.Println(movie1)
			fmt.Println(string(b))
			conn.Do("LPUSH", "movies", string(b))
		}
	}
}

func main() {

	var aCustomer customer.Customer
	conn := aCustomer.GetRedisConn()
	conn.Do("select", 2)
	defer conn.Close()
	//从redis中取出电影
	var moviesList []movie.BaseMovie
	// movies, _ := redis.Strings(conn.Do("LRANGE", "movies", 0, -1))
	// fmt.Println(movies)
	// for _, value := range movies {
	// 	tempMovie := movie.BaseMovie{}
	// 	err := json.Unmarshal([]byte(value), &tempMovie)
	// 	if err == nil {
	// 		moviesList = append(moviesList, tempMovie)
	// 	}
	// }
	movies, _ := redis.Values(conn.Do("LRANGE", "movies", 0, -1))
	//fmt.Println(movies)
	for _, value := range movies {
		tempMovie := movie.BaseMovie{}
		err := json.Unmarshal([]byte(value.([]uint8)), &tempMovie)
		if err == nil {
			moviesList = append(moviesList, tempMovie)
		}
	}
	//fmt.Println(moviesList)

	//初始化
	var rentalList []rental.Rental
	movie1 := movie.RegularMovie{moviesList[0]}
	// movie2 := movie.NewIssueMovie{movie.BaseMovie{Title: "The Godfather Part II", TypeCode: movie.RELEASE}}
	// movie3 := movie.ChildrenMovie{movie.BaseMovie{Title: "Casablanca", TypeCode: movie.CHILDS}}
	movie2 := movie.RegularMovie{moviesList[1]}
	movie3 := movie.RegularMovie{moviesList[2]}
	rentalList = append(rentalList, rental.Rental{movie1, 3})
	rentalList = append(rentalList, rental.Rental{movie2, 1})
	rentalList = append(rentalList, rental.Rental{movie3, 7})
	customer := customer.Customer{"Tom", rentalList}
	result := customer.Statement()
	fmt.Println(result)
}
