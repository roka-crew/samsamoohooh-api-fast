package domain

import "time"

type Group struct {
	ID               int `storm:"id,increment"`
	BookTitle        string
	BookAuthor       string
	BookPageMax      int
	BookPageCount    int
	BookPublisher    *string
	BookIntroduction *string

	Users  []UserMembership
	Topics []Topic

	CreatedAt time.Time
	UpdatedAt time.Time
}

type GroupMembership struct {
	GroupID  int
	JoinedAt time.Time
}
