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
