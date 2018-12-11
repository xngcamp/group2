package movie

type MovieKind interface {
	GetMoney(daysRented int) float64
	GetPoints(daysRented int) int
}
