package domain

import "time"

type Goal struct {
	Deadline  time.Time
	PageRange int
	Topics    []Topic

	CreatedAt time.Time
	UpdatedAt time.Time
}
