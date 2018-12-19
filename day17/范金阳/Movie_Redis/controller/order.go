package controller

import "Movie_Redis/api"

type Order struct {
	Name  string
	Rente []*api.Rente

}

func NewOrder() *Order {
	return &Order{}
}

func (order *Order) AddRente(rente *api.Rente)  {
		order.Rente =append(order.Rente,rente)
}
