package model

import "time"

type Match struct {
	ID       int64
	GameId   int64
	Notes    string
	DateTime time.Time
	Location string
	// players
}
