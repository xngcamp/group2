package movie



//普通电影
type OrdinaryMovie struct {
	Title     string //片名
}

func (m *OrdinaryMovie) PutTitle(s string)()  {
	m.Title=s
}


func (m *OrdinaryMovie) GetTitle() string {
	return m.Title
}
//获取价格
func (ordMovie *OrdinaryMovie) GetCharge(daysRented int) float64 {
	result :=0.0
	result += 2
	if daysRented > 2 {
		result += float64(daysRented-2.0) * float64(1.5)
	}
	return result
}
//获取积分
func (ordMovie *OrdinaryMovie) GetFrequentRenterPoints(daysRented int) int  {
	return 1
}


