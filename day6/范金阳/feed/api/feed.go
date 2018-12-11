package api
type Feed struct {
	ID int `json:"id" bson:"_id"`
	TXT string `json:"txt" bson:"txt"`
}

func NewFeed() *Feed {
	return &Feed{}
}
