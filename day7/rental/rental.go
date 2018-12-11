package rental

import "github.com/xngcamp/group2/day7/movie"

type Rental struct {
	Movie movie.Movie
	Days int
}

func (r *Rental) GetPrice() float64 {
	return r.Movie.GetPrice(r.Days)
}

func (r *Rental) GetPoints() int {
	return r.Movie.GetPoints(r.Days)
}