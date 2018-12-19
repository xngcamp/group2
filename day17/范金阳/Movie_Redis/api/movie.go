package api

const(
	EXM =iota
	NEW
	CHILD
	TRAD
)


type Movie struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Year string `json:"year"`
	Author string `json:"author"`
	Star string `json:"star"`
	PriceCode int `json:"price_code"`
}
