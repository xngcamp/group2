package controller

import (
	"Movie_Redis/service"
	"fmt"
)

func(order *Order) CalOrder()  {
	charge , points := 0.0, 0
	movie := service.NewMovie()

	fmt.Printf("%v 租赁记录：\n",order.Name)
	for i,r := range order.Rente {
		c , p :=0.0,0
		movie.Id=r.Id
		p=movie.GetFrequentRenterPoints(r.Day)
		points+=p
		c,title:=movie.GetCharge(r.Day)
		charge+=c
		fmt.Printf("	  %v. %v %v元 %v积分\n",i+1,title,c,p)
	}
	fmt.Printf("共计 %v元 %v积分\n",charge,points)



}
