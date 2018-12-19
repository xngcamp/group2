package movie

type BaseMovie struct {
	Id       int     `json:"id"`
	Title    string  `json:"title"`  //片名
	Year     int64   `json:"year"`   //年份
	Author   string  `json:"author"` //作者
	Star     float64 `json:"star"`   //评分
	TypeCode int     `json="typecode"`
}
