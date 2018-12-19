package controller

import (
	"Movie_Redis/api"
	"Movie_Redis/model"
	"fmt"

)

func(moive *Order) AddMovie()  {
	mov :=model.NewMovie()
	mov.Id=1
	mov.Title="乘风破浪"
	mov.Year="2017-5-6"
	mov.Author="李安"
	mov.Star="五星"
	mov.PriceCode=api.NEW
	ok1 :=mov.Add()


	mov2 :=model.NewMovie()
	mov2.Id=2
	mov2.Title="冰河世纪"
	mov2.Year="2017-8-9"
	mov2.Author="张艺谋"
	mov2.Star="四星"
	mov2.PriceCode=api.CHILD
	ok2 :=mov2.Add()


	mov3 :=model.NewMovie()
	mov3.Id=3
	mov3.Title="集结号"
	mov3.Year="2001-12-1"
	mov3.Author="孙策"
	mov3.Star="五星"
	mov3.PriceCode=api.TRAD
	ok3:=mov3.Add()
	if ok1&&ok2&&ok3{
		fmt.Println("success")
		return
	}
	fmt.Println("error")
}
