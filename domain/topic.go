package domain

import (
	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	Title   string
	Content string

	GoalID int
	UserID int
}

type Topics []Topic

func (t Topics) IsEmpty() bool {
	return len(t) == 0
}

func (t Topics) First() Topic {
	if t.IsEmpty() {
		return Topic{}
	}
	return t[0]
}
