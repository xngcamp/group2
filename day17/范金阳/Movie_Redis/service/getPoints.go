package service

import (
	"Movie_Redis/api"
	"Movie_Redis/model"
)

func (m *Movie) GetFrequentRenterPoints(daysRented int) int {
	movie := model.NewMovie()
	movie.Id=m.Id
	apiMovie := movie.Get()
	if apiMovie.PriceCode == api.NEW && daysRented > 1 {
		return 2
	} else {
		return 1
	}
}
