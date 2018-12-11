package movie



//新电影
type NewMovie struct {
	Title     string //片名
}
func (m *NewMovie) PutTitle(s string)()  {
	m.Title=s
}


func (m *NewMovie) GetTitle() string {
	return m.Title
}
//获取价格
func (newMovie *NewMovie)GetCharge(daysRented int) float64 {
	result :=0.0
	result += float64(daysRented) * float64(3)
	return result
}

//获取积分
func (newMovie *NewMovie) GetFrequentRenterPoints(daysRented int) int  {
	if daysRented > 1 {
		return 2
	} else {
		return 1
	}
}



