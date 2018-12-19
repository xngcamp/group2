package model

import "Movie_Redis/api"

type Movie api.Movie

func NewMovie() *Movie {
	return &Movie{}
}
