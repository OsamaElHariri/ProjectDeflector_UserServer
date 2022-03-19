package users

type GameStats struct {
	Games int
	Wins  int
}

type User struct {
	Id        string
	Uuid      string
	Nickname  string
	Color     string
	GameStats GameStats
}
