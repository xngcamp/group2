/*
Package user just for demo
*/
package user

import (
	"camp/user/mongo"

	"github.com/globalsign/mgo/bson"

	mgo "github.com/globalsign/mgo"
)

// UserModel 定义db对应的类型
type UserModel struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Nick     string        `json:"nick" bson:"_nick"`         //用户名
	Sex      byte          `json:"sex" bson:"_sex"`           //性别 1|2
	Email    string        `json:"email" bson:"_email"`       //电子邮箱
	Password string        `json:"password" bson:"_password"` //密码
}

// Db 返回db name
func (userModel *UserModel) Db() (db string) {
	return "user"
}

// Table 返回table name
func (userModel *UserModel) Table() (table string) {
	return "user"
}

// GetC 返回db col
func (userModel *UserModel) GetC() (c *mgo.Collection) {
	db, table := userModel.Db(), userModel.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}
