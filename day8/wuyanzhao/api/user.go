package api

//实体类
type UserApi struct {
	ID       string `json:"id" bson:"_id"`
	Nick     string `json:"nick" bson:"_nick"`         //用户名
	Sex      byte   `json:"sex" bson:"_sex"`           //性别 1|2
	Email    string `json:"email" bson:"_email"`       //电子邮箱
	Password string `json:"password" bson:"_password"` //密码
}

//生成User对象
func NewUserApi() *UserApi {
	return &UserApi{}
}
