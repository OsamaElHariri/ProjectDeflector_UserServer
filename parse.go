package main

import "projectdeflector.users/users"

type User struct {
	Uuid     string `json:"uuid"`
	Nickname string `json:"nickname"`
	Color    string `json:"color"`
}

func parseUser(user users.User) User {
	return User{
		Uuid:     user.Uuid,
		Nickname: user.Nickname,
		Color:    user.Color,
	}
}
