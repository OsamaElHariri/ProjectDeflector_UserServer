package main

import "projectdeflector.users/users"

type User struct {
	Id        string    `json:"id"`
	Nickname  string    `json:"nickname"`
	Color     string    `json:"color"`
	GameStats GameStats `json:"gameStats"`
}

type GameStats struct {
	Games int `json:"games"`
	Wins  int `json:"wins"`
}

func parseUser(user users.User) User {
	return User{
		Id:       user.Id,
		Nickname: user.Nickname,
		Color:    user.Color,
		GameStats: GameStats{
			Games: user.GameStats.Games,
			Wins:  user.GameStats.Wins,
		},
	}
}
