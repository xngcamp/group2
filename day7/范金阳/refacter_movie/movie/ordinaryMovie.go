package movie



//普通电影
type OrdinaryMovie struct {
	Movie
}

//获取价格
func (ordMovie *OrdinaryMovie) GetMoney(daysRented int) float64 {
	result :=0.0
	result += 2
	if daysRented > 2 {
		result += float64(daysRented-2.0) * float64(1.5)
	}
	return result
}
//获取积分
func (ordMovie *OrdinaryMovie) GetPoints(daysRented int) int  {
	return 1
}


