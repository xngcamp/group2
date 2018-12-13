package api

//实体类
type LoginApi struct {
	//ID        string `json:"id" bson:"_id"`
	Email    string `json:"email" bson:"_email"`       //电子邮箱
	Password string `json:"password" bson:"_password"` //密码
	//Usertoken string `json:"usertoken" bson:"_usertoken"`
}

//生成User对象
func NewLoginApi() *LoginApi {
	return &LoginApi{}
}
