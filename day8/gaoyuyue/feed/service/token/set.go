package token

import (
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/model"
	"gopkg.in/mgo.v2/bson"
)

func (t *Token) Set(tokenApi *api.Token) (err error) {
	tokenModel := model.NewToken()
	tokenApi.Id = bson.NewObjectId()
	tokenModel.UserId = tokenApi.UserId
	tokenModel.Id = tokenApi.Id
	err = tokenModel.Set()
	if err != nil {
		return
	}
	return
}
