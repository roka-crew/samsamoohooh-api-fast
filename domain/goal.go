package domain

import (
	"time"

	"gorm.io/gorm"
)

type Goal struct {
	gorm.Model
	Deadline  time.Time
	PageRange int

	GroupID uint
	Topics  []Topic
}

type Goals []Goal

func (g Goals) First() Goal {
	if len(g) == 0 {
		return Goal{}
	}
	return g[0]
}

func (g Goals) IsEmpty() bool {
	return len(g) == 0
}
