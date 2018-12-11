package main

import (
	"github.com/xngcamp/group2/day7/customer"
	"github.com/xngcamp/group2/day7/movie"
	"github.com/xngcamp/group2/day7/rental"
)

func main() {
	c:= &customer.Customer{Name: "Tom"}
	commonMovie := &movie.CommonMovie{MovieCommon:movie.MovieCommon{Name:"Star Wars"}}
	r1 := &rental.Rental{commonMovie, 3}
	c.AddRental(r1)
	newReleaseMovie := &movie.NewReleaseMovie{MovieCommon:movie.MovieCommon{Name:"The Godfather Part II"}}
	r2 := &rental.Rental{newReleaseMovie, 2}
	c.AddRental(r2)
	childrenMovie := &movie.ChildrenMovie{MovieCommon:movie.MovieCommon{Name:"Casablanca"}}
	r3 := &rental.Rental{childrenMovie, 7}
	c.AddRental(r3)
	c.Print()
}
