package movie



//儿童电影
type ChildMovie struct {
	Mov Movie
}


//获取价格
func (childMovie *ChildMovie) GetMoney(daysRented int) float64 {
	result :=0.0
	result += 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * float64(1.5)
	}
	return result
}
//获取积分
func (childMovie *ChildMovie) GetPoints(daysRented int) int  {
	return 1
}