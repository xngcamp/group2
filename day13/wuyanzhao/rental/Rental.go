package rental

import "camp/week2/day2/wuyanzhao/movie"

type Rental struct {
	Movie movie.MovieInterface
	RentDays int
}

func (r Rental) GetCharge() float64 {
	return r.Movie.GetCost(r.RentDays)
}

func (r Rental) GetFrequentRenterPoints() int {
	return r.Movie.GetPoint(r.RentDays)
}
