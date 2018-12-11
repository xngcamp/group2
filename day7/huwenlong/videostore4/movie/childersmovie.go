package movie

type ChildrenMovie struct {
	Title string
}

func (movie ChildrenMovie) GetTitle() string {
	return movie.Title
}

func (movie ChildrenMovie) GetPriceCode() int {
	return 0
}

func (movie ChildrenMovie) GetCharge(daysRented int) float64 {
	result := 0.0
	result += 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * float64(1.5)
	}
	return result
}

func (movie ChildrenMovie) GetFrequentRenterPoints(daysRented int) int {
	return 1
}

