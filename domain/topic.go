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
