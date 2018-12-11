package movie

type NewMovie struct {
	Title string
}

func (movie NewMovie) GetTitle() string {
	return movie.Title
}

func (movie NewMovie) GetPriceCode() int {
	return 0
}
func (movie NewMovie) GetCharge(daysRented int) float64 {
	result := 0.0
	result += float64(daysRented) * float64(3)
	return result
}

func (movie NewMovie) GetFrequentRenterPoints(daysRented int) int {
	if daysRented > 1 {
		return 2
	}else {
		return 1
	}
}
