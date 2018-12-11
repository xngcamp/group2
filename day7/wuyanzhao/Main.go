package main

import (
	"camp/week2/day2/wuyanzhao/customer"
	"camp/week2/day2/wuyanzhao/movie"
	"camp/week2/day2/wuyanzhao/rental"
	"fmt"
)

func main(){
	var rentalList []rental.Rental
	movie1:=movie.RegularMovie{movie.BaseMovie{Title: "The Godfather Part II", TypeCode: movie.NORMAL}}
	movie2 := movie.NewIssueMovie{movie.BaseMovie{Title: "The Godfather Part II", TypeCode: movie.RELEASE}}
	movie3 := movie.ChildrenMovie{movie.BaseMovie{Title: "Casablanca", TypeCode: movie.CHILDS}}
	rentalList=append(rentalList,rental.Rental{movie1,3})
	rentalList=append(rentalList,rental.Rental{movie2,1})
	rentalList=append(rentalList,rental.Rental{movie3,7})
	customer:=customer.Customer{"Tom",rentalList}
	result := customer.Statement()
	fmt.Println(result)
}

