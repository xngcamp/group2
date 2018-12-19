package movie

const (
	REGULAR  = iota
	NEW_RELEASE
	CHILDREN
)
type Movie struct {
	PriceCode int
	Title string
}

func (movie Movie) GetTitle() string {
	return movie.Title
}

func (movie Movie) GetPriceCode() int {
	return movie.PriceCode
}



