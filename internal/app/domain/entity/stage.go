package entity

import "time"

type StageWeight struct {
	ID         int64
	Stage      string
	Weight     float64
	OrderIndex int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
