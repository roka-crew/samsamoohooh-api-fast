package presenter

import "github.com/roka-crew/domain"

type CreateGroupParams = domain.Group

type ListGroupsParams struct {
	IDs       []uint
	Limit     int
	WithUsers bool
	WithGoals bool
}

type AddGroupUsersParams struct {
	GroupID uint
	UserIDs []uint
}

type ListGroupUsersParams struct {
	GroupID uint
	UserIDs []uint
}
