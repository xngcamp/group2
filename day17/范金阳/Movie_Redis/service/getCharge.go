package service

import (
	"Movie_Redis/api"
	"Movie_Redis/model"
)

func (m *Movie) GetCharge(daysRented int) (float64,string) {
	result := 0.0
	movie := model.NewMovie()
	movie.Id=m.Id
	apiMovie := movie.Get()
	name :=apiMovie.Title
	switch apiMovie.PriceCode {
	case api.NEW:
		result += 2
		if daysRented > 2 {
			result += float64(daysRented-2.0) * float64(1.5)
		}
	case api.CHILD:
		result += float64(daysRented) * float64(3)
	case api.TRAD:
		result += 1.5
		if daysRented > 3 {
			result += float64(daysRented-3) * float64(1.5)
		}
	}
	return result,name
}
