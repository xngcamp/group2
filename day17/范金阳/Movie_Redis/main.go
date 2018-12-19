package main

import (
	"Movie_Redis/api"
	"Movie_Redis/controller"
)

func main()  {
	movie :=controller.NewOrder()
	movie.AddMovie()

	order := controller.NewOrder()
	order.Name="tom"
	rente1 := api.Rente{1,1}
	rente2 := api.Rente{2,2}
	rente3 := api.Rente{3,2}
	order.AddRente(&rente1)
	order.AddRente(&rente2)
	order.AddRente(&rente3)
	order.CalOrder()
}
