package movie

import (
	"encoding/json"
	redis2 "github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2/bson"
	"gothinking/refactoring/videostore4/redis"
)

const (
	REGULAR     = iota //0
	NEW_RELEASE        //1
	CHILDREN           // 2
)

type Movie struct {
	Id        bson.ObjectId   `json:"id"`
	Title     string          `json:"title"` //片名
	PriceCode int             `json:"price_code"` //价格代号
}

func (m *Movie) Set() (err error) {
	m.Id = bson.NewObjectId()
	conn := redis.GetRedisConn()
	bs, err := json.Marshal(m)
	if err != nil {
		return
	}
	_, err = conn.Do("SET", m.Id, bs)
	if err != nil {
		return
	}
	return
}

func (m *Movie) Get() (err error) {
	conn := redis.GetRedisConn()
	bs, err := redis2.Bytes(conn.Do("GET", m.Id))
	if err != nil {
		return
	}
	err = json.Unmarshal(bs, m)
	if err != nil {
		return
	}
	return
}

func (m Movie) GetTitle() string {
	return m.Title
}

func (m Movie) GetPriceCode() int {
	return m.PriceCode
}

func (m Movie) GetCharge(daysRented int) float64 {
	result := 0.0
	switch m.GetPriceCode() {
	case REGULAR:
		result += 2
		if daysRented > 2 {
			result += float64(daysRented-2.0) * 1.5
		}
	case NEW_RELEASE:
		result += float64(daysRented) * 3
	case CHILDREN:
		result += 1.5
		if daysRented > 3 {
			result += float64(daysRented-3) * 1.5
		}
	}
	return result
}

func (m Movie) GetFrequentRenterPoints(daysRented int) int {
	if m.GetPriceCode() == NEW_RELEASE && daysRented > 1 {
		return 2
	} else {
		return 1
	}
}
