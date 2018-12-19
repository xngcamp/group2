package customer

import (
	"fmt"
	"videostore4/movie"
	"videostore4/rental"
)

/*
重构：
1. 把计算每一个租金的方法充Statement中抽取出来，形成新的amountFor()方法
2. amountFor() 参数更合理的命名
3. amountFor()中计算租金，只用了了rental，并没有用到customer， 放在customer中不合理，继续抽取
3. struct中作用域, 不暴露内部信息
*/

type Customer struct {
	name    string
	rentals []rental.Rental
}

func (c *Customer) Init(name string) {
	c.name = name
}

func (c *Customer) AddRental(arg rental.Rental) {
	c.rentals = append(c.rentals, arg)
}

func (c Customer) GetName() string {
	return c.name
}
func (c Customer) GetRentals() []rental.Rental {
	return c.rentals
}
func (c Customer) Statement() string {
	var totalAmount float64      //消费总额
	var frequentRenterPoints int //常客积分
	var result= "Rental Record for " + c.GetName() + ":\n"
	for index, r := range c.GetRentals() {

		thisAmount := 0.0
		// determine amounts for each record ,  各种片子价格不同
		aMovie := r.GetMovie()
		switch aMovie.PriceCode {
		case movie.REGULAR:
			thisAmount += 2
			if r.GetDaysRented() > 2 {
				thisAmount += float64(r.GetDaysRented()-2) * 1.5
			}
		case movie.NEW_RELEASE:
			thisAmount += float64(r.GetDaysRented()) * 2
		case movie.CHILDREN:
			thisAmount += 1.5
			if r.GetDaysRented() > 3 {
				thisAmount += float64(r.GetDaysRented()-3) * 1.5
			}
		}

		// add frequent renter points (累计常客积分)
		frequentRenterPoints++

		// add bonus for a two day new release rental
		if aMovie.PriceCode == movie.NEW_RELEASE && r.GetDaysRented() > 1 {
			frequentRenterPoints++
		}

		//show figures for this rental (显示此次租赁的数据)
		thisResult := fmt.Sprintf("\t%d. %s\t%.2f\n", index+1, aMovie.Title, thisAmount)
		result += thisResult
		totalAmount += thisAmount
	}
	return result
}