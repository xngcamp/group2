/*
Package user just for demo
*/
package user

import (
	"camp/user/mongo"

	mgo "github.com/globalsign/mgo"
)

// LoginModel 定义db对应的类型
type LoginModel struct {
	Usertoken string `json:"usertoken" bson:"_usertoken"`
	UserID    string `json:"userid" bson:"_userid"`
}

// Db 返回db name
func (loginModel *LoginModel) Db() (db string) {
	return "user"
}

// Table 返回table name
func (loginModel *LoginModel) Table() (table string) {
	return "login"
}

// GetC 返回db col
func (loginModel *LoginModel) GetC() (c *mgo.Collection) {
	db, table := loginModel.Db(), loginModel.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}
