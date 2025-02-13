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

type Topic struct {
	UserID  int
	Title   string
	Content string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Goal struct {
	Deadline  time.Time
	PageRange int
	Topics    []Topic

	CreatedAt time.Time
	UpdatedAt time.Time
}

type GroupMembership struct {
	GroupID  int
	JoinedAt time.Time
}

type UserMembership struct {
	UserID   int
	JoinedAt time.Time
}
