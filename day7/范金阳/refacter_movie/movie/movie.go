package movie

const (
	REGULAR     = iota //0
	NEW_RELEASE        //1
	CHILDRES           // 2
)

type Movie interface {
	GetCharge(daysRented int) float64
	GetFrequentRenterPoints(daysRented int) int
	GetTitle() string
	PutTitle(s string)

	//
	//Title     string //片名
	//PriceCode int    //价格代号
}

//func (m Movie) GetTitle() string {
//	return m.Title
//}
//
//func (m Movie) GetPriceCode() int {
//	return m.PriceCode
//}

//func (m Movie) GetCharge(daysRented int) float64 {
//	result := 0.0
//	switch m.GetPriceCode() {
//	case REGULAR:
//		result += 2
//		if daysRented > 2 {
//			result += float64(daysRented-2.0) * float64(1.5)
//		}
//	case NEW_RELEASE:
//		result += float64(daysRented) * float64(3)
//	case CHILDRES:
//		result += 1.5
//		if daysRented > 3 {
//			result += float64(daysRented-3) * float64(1.5)
//		}
//	}
//	return result
//}
//
//func (m Movie) GetFrequentRenterPoints(daysRented int) int {
//	if m.GetPriceCode() == NEW_RELEASE && daysRented > 1 {
//		return 2
//	} else {
//		return 1
//	}
//}


//
////获取消费价格
//func (m Movie) GetCharge(daysRented int) float64 {
//	result := 0.0
//	var kind  MovieKind
//	switch m.GetPriceCode() {
//	case REGULAR:
//		kind = new(OrdinaryMovie)
//	case NEW_RELEASE:
//		kind = new(NewMovie)
//	case CHILDRES:
//		kind = new(ChildMovie)
//	}
//	result += kind.GetMoney(daysRented)
//	return result
//}
//
////获取积分
//func (m Movie) GetFrequentRenterPoints(daysRented int) int {
//	var kind  MovieKind
//	switch m.GetPriceCode() {
//	case REGULAR:
//		kind = new(OrdinaryMovie)
//	case NEW_RELEASE:
//		kind = new(NewMovie)
//	case CHILDRES:
//		kind = new(ChildMovie)
//	}
//	return kind.GetPoints(daysRented)
//}