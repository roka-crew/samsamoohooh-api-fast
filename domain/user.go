package domain

import "time"

type User struct {
	ID         int `storm:"id,increment"`
	Nickname   string
	Resolution *string

	Groups []GroupMembership

	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserMembership struct {
	UserID   int
	JoinedAt time.Time
}
