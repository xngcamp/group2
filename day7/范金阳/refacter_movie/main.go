package main

import (
	"fmt"
	"refacter_movie/customer"
	"refacter_movie/movie"
	"refacter_movie/rental"
)

/*

作业总结：
	在movie包下新建了三个结构体继承自Movie，
同时新建了一个movieInterface的接口，其中包含GetMoney获取价钱的方法和getPoint获取积分的方法，
并将三个结构体实现结构中的方法。

关于新增一个电影类型的思考：
	可以直接在movie包中新建一个实体类,并且实现movieInterface接口，并在movie中添加相应
的代表编号


*/


func main() {
	//aCustomer := new(customer.Customer)
	var aCustomer customer.Customer
	aCustomer.Init("Tom")

	var rental1 rental.Rental
	//m1 := movie.OrdinaryMovie{movie.Movie{Title: "Star Wars", PriceCode: movie.REGULAR}}
	//movie1 := movie.Movie{Title: "Star Wars", PriceCode: movie.REGULAR}
	//var movie1 movie.Movie
	movie1 :=new(movie.OrdinaryMovie)
	movie1.PutTitle("Star Wars")
	//movie1 := m1.Movie
	rental1.Init(movie1, 3)
	aCustomer.AddRental(rental1)

	var rental2 rental.Rental
	//movie2 := movie.Movie{Title: "The Godfather Part II", PriceCode: movie.NEW_RELEASE}
	movie2 := new(movie.NewMovie)
	movie2.PutTitle("The Godfather Part II")
	rental2.Init(movie2, 1)
	aCustomer.AddRental(rental2)

	var rental3 rental.Rental
	movie3 := new(movie.ChildMovie)
	movie3.PutTitle("Casablanca")
	rental3.Init(movie3, 7)
	aCustomer.AddRental(rental3)

	result := aCustomer.Statement()
	fmt.Println(result)
}
