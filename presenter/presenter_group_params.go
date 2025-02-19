package presenter

import "github.com/roka-crew/domain"

type CreateGroupParams = domain.Group

type ListGroupsParams struct {
	IDs       []uint
	Limit     int
	WithUsers bool
	WithGoals bool
}

type PatchGroupParams struct {
	// condition
	GroupID uint

	// change
	BookTitle        *string
	BookAuthor       *string
	BookPageMax      *int
	BookPageCount    *int
	BookPublisher    *string
	BookIntroduction *string
}

type DeleteGroupParams struct {
	GroupID uint
}

type AddGroupUsersParams struct {
	GroupID uint
	UserIDs []uint
}

type ListGroupUsersParams struct {
	GroupID uint
	UserIDs []uint
}
