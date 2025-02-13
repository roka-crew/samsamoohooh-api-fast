package domain

import "time"

type Topic struct {
	UserID  int
	Title   string
	Content string

	CreatedAt time.Time
	UpdatedAt time.Time
}
