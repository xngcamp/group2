package movie



//新电影
type NewMovie struct {
	Mov Movie
}

//获取价格
func (newMovie *NewMovie) GetMoney(daysRented int) float64 {
	result :=0.0
	result += float64(daysRented) * float64(3)
	return result
}

//获取积分
func (newMovie *NewMovie) GetPoints(daysRented int) int  {
	if daysRented > 1 {
		return 2
	} else {
		return 1
	}
}



