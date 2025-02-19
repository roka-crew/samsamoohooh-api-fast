package presenter

import "github.com/roka-crew/domain"

type CreateUserParams = domain.User

type ListUsersParams struct {
	IDs        []uint
	Nicknames  []string
	WithGroups bool
	WithTopics bool
	Limit      int
}

type PatchUserParams struct {
	// condition
	UserID *uint

	// change
	Nickname   *string
	Resolution *string
}

type DeleteUserParams struct {
	UserID         uint
	Nickname       string
	WithHardDelete bool
}

type ListUserGroupsParams struct {
	UserID uint
}

type RemoveGroupUserParams struct {
	UserID  uint
	GroupID uint
}

type ExistsGroupUserParams struct {
	GroupID uint
	UserID  uint
}
