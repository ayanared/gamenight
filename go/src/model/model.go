package model

import "time"

type Player struct {
	Name  string
	Email string
}

type Game struct {
	ID         string
	Time       time.Time
	Name       string
	BGGLink    string
	Location   string
	MinPlayers int
	MaxPlayers int
}

type GamePlayer struct {
	PlayerEmail string
	GameID      string
}
