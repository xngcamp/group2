package customer

import (
	"fmt"
	"github.com/xngcamp/group2/day7/gaoyuyue/rental"
)

type Customer struct {
	Name string
	Rentals []*rental.Rental
}

func (c *Customer) AddRental(rental *rental.Rental)  {
	c.Rentals = append(c.Rentals, rental)
}

func (c Customer) GetTotalPoints() int {
	result := 0
	for _, r := range c.Rentals {
		result += r.GetPoints()
	}
	return result
}

func (c Customer) GetTotalPrice() float64 {
	result := 0.0
	for _, r := range c.Rentals {
		result += r.GetPrice()
	}
	return result
}

func (c Customer) Print() {
	fmt.Printf("Rental Record for %v \n", c.Name)
	for i, r := range c.Rentals{
		fmt.Printf("\t %d .%v  %.2f\n", i+1, r.Movie.GetName(), r.GetPrice())
	}
	fmt.Printf("Amount owned is %.2f\n", c.GetTotalPrice())
	fmt.Printf("%v earned %d frequent renter points.\n", c.Name, c.GetTotalPoints())
}