package service

import (
	"camp/feed/service/feed"
	"camp/feed/service/user"
)

func NewFeed() *feed.Feed{
	return &feed.Feed{}
}

func NewUser() *user.User {
	return &user.User{}
}


