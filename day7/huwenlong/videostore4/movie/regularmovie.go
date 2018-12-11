package movie

type RegularMovie struct {
	Title string
}

func (movie RegularMovie) GetTitle() string {
	return movie.Title
}

func (movie RegularMovie) GetPriceCode() int {
	return 0
}
func (movie RegularMovie) GetCharge(daysRented int) float64 {
	result := 0.0
	result += 2
	if daysRented > 2 {
		result += float64(daysRented-2.0) * float64(1.5)
	}
	return result
}

func (movie RegularMovie) GetFrequentRenterPoints(daysRented int) int {
	return 1
}