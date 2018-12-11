package movie

//每个电影结构体去实现这个接口
type Movie interface {
	//定义movie接口中的虚函数
	GetTitle() string
	GetPriceCode() int
	GetCharge(daysRented int) float64
	GetFrequentRenterPoints(daysRented int) int
}
