package model

import "github.com/ericlagergren/decimal"

type Score struct {
	ID         int64
	CategoryID int64
	PlayerID   int64
	Value      decimal.Big
}
