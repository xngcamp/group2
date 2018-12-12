package model

import (
	"camp/feed/model/feed"
	"camp/feed/model/user"
)

func NewFeed() *feed.Feed {
	return &feed.Feed{}
}

func NewUser() *user.User {
	return &user.User{}
}