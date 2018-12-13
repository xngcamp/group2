package session

import (
	"errors"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/api"
	"github.com/xngcamp/group2/day8/gaoyuyue/feed/model"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func (s *Session) GetById(id string) (sessionApi *api.Session, err error) {
	sessionApi = api.NewSession()
	sessionModel := model.NewSession()
	sessionModel.Id = bson.ObjectIdHex(id)
	err = sessionModel.Get()
	if err != nil {
		return
	}
	if time.Now().UTC().After(sessionModel.Timeout) {
		err = errors.New("session timeout")
		if e := sessionModel.Del(); e != nil {
			return sessionApi, e
		}
		return
	}

	duration, err := time.ParseDuration("30m")
	if err != nil {
		return
	}
	sessionModel.Timeout = time.Now().Add(duration)
	err = sessionModel.Update()
	if err != nil {
		return
	}
	sessionApi.Id = sessionModel.Id
	sessionApi.UserId = sessionModel.UserId
	sessionApi.Timeout = sessionModel.Timeout
	return
}