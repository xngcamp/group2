package rental

import "videostore4/movie"

type Rental struct {
	aMovie     movie.Movie
	daysRented int //租期
}

func (r *Rental) Init(movie movie.Movie, daysRented int) {
	r.aMovie = movie
	r.daysRented = daysRented
}

func (r Rental) GetDaysRented() int {
	return r.daysRented
}

func (r Rental) GetMovie() movie.Movie {
	return r.aMovie
}
