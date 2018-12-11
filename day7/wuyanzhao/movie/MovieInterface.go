package movie

type MovieInterface interface {
	GetCost(days int) float64
	GetPoint(days int) int
	GetTitle() string
}