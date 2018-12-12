package session

import (
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/model"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func (s *Session) Add(userId bson.ObjectId) (sessionApi *api.Session,err error) {
	sessionApi = api.NewSession()
	duration, e := time.ParseDuration("30m")
	if e != nil {
		return sessionApi, e
	}
	sessionApi.Timeout = time.Now().Add(duration)
	sessionApi.Id = bson.NewObjectId()
	sessionApi.UserId = userId
	sessionModel := model.NewSession()
	sessionModel.Id = sessionApi.Id
	sessionModel.Timeout = sessionApi.Timeout
	sessionModel.UserId = sessionApi.UserId
	if err = sessionModel.Add(); err != nil {
		return
	}
	return
}
