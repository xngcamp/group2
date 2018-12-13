/*
Package user just for demo
*/
package user

import (
	"camp/user/api"
	"camp/user/mongo"

	mgo "github.com/globalsign/mgo"
)

// UserModel 定义db对应的类型
type UserModel api.UserApi

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
