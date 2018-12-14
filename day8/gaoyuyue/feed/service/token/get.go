package token

import (
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/model"
	"gopkg.in/mgo.v2/bson"
)

func (t *Token) GetById(id string) (tokenApi *api.Token, err error) {
	tokenModel := model.NewToken()
	tokenModel.Id = bson.ObjectIdHex(id)
	err = tokenModel.Get()
	if err != nil {
		return
	}
	tokenApi = api.NewToken()
	tokenApi.Id = tokenModel.Id
	tokenApi.UserId = tokenModel.UserId
	return
}
