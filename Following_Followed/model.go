package main

import "time"

type User struct {
	Email    string
	UserName string `datastore:"-"`
	Password string `json:"-"`
	Following []string `datastore:"-"`
	FollowedBy []string `datastore:"-"`
}

type SessionData struct {
	User
	LoggedIn  bool
	LoginFail bool
	Tweets []Tweet
	ViewingUser User
	FollowingUser bool
}

type Tweet struct {
	Msg string
	Time time.Time
	UserName string
}

type Follow struct {
	Following string
	Follower string
