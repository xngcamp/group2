package movie

const (
	NORMAL=iota
	RELEASE
	CHILDS
)
type Movie struct {
	BaseMovie
}

//func (m Movie)GetCharge(rentDays int) float64{
//	result:=0.0
//	switch m.TypeCode{
//	case NORMAL:
//		result += 2
//		if rentDays > 2 {
//			result += float64(rentDays-2.0) * float64(1.5)
//		}
//	case RELEASE:
//		result += float64(rentDays) * float64(3)
//	case CHILDS:
//		result += 1.5
//		if rentDays > 3 {
//			result += float64(rentDays-3) * float64(1.5)
//		}
//	}
//	return result
//}
//
//func (m Movie)GetFrequentRentPoints(rentDays int) int{
//	if m.TypeCode == RELEASE && rentDays > 1 {
//		return 2
//	} else {
//		return 1
//	}
//}