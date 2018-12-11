package api

//创建feed结构体
type Feed struct {
	Txt string `json:"txt" bson:"_txt"`
	Id int64 `json:"id" bson:"_id"`
}

//暴露feed结构体 生成feed
func NewFeed()  *Feed{
	return &Feed{}
}