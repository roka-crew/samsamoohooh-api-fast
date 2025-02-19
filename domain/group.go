package domain

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	BookTitle        string
	BookAuthor       string
	BookPageMax      int
	BookPageCount    int
	BookPublisher    *string
	BookIntroduction *string

	Goals []Goal
	Users []User `gorm:"many2many:user_group_mapper;"`
}

type Groups []Group

func (g Groups) IsEmpty() bool {
	return len(g) == 0
}

func (g Groups) First() Group {
	if g.IsEmpty() {
		return Group{}
	}

	return g[0]
}
