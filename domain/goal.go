package domain

import (
	"time"

	"gorm.io/gorm"
)

type Goal struct {
	gorm.Model
	Deadline  time.Time
	PageRange int

	GroupID int
	Topcis  []Topic
}
