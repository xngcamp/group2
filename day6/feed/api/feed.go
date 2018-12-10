package api

//实体类
type Feed struct {
	ID  string `json:"id" bson:"_id"`
	Txt string `json:"txt" bson:"_txt"` //发布内容
}

//生成Feed对象
func NewFeed() *Feed {
	return &Feed{}
}
