package movie



//儿童电影
type ChildMovie struct {
	Title     string //片名
}

func (m *ChildMovie) PutTitle(s string)()  {
	m.Title=s
}

func (m *ChildMovie) GetTitle() string {
	return m.Title
}
//获取价格
func (childMovie *ChildMovie) GetCharge(daysRented int) float64 {
	result :=0.0
	result += 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * float64(1.5)
	}
	return result
}
//获取积分
func (childMovie *ChildMovie) GetFrequentRenterPoints(daysRented int) int  {
	return 1
}