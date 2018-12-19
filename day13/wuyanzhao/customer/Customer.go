package customer

import (
	"camp/week2/day2/wuyanzhao/redis"
	"camp/week2/day2/wuyanzhao/rental"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type Customer struct {
	Name     string
	RentList []rental.Rental
}

//获取redis连接
func (c *Customer) GetRedisConn() (conn redis.Conn) {
	r := myredis.Pool
	conn = r.Get()
	return
}

func (c Customer) Statement() string {
	result := "Rental Record for " + c.Name + ":\n"
	var totalAmount float64      //消费总额
	var frequentRenterPoints int //常客积分
	for index, value := range c.RentList {
		thisAmount := value.GetCharge()
		totalAmount += thisAmount
		frequentRenterPoints += value.GetFrequentRenterPoints()
		//show figures for this rental (显示此次租赁的数据)
		thisResult := fmt.Sprintf("\t%d. %s\t%.2f\n", index+1, value.Movie.GetTitle(), thisAmount)
		result += thisResult
	}
	// add footer lines (添加页脚信息)
	result += fmt.Sprintf("Amount owned is %.2f\n", totalAmount)
	result += fmt.Sprintf("%s earned %d frequent renter points.\n", c.Name, frequentRenterPoints)
	return result
}
