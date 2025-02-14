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
	UserID     *uint
	Nickname   *string
	Resolution *string
}

type DeleteUserParams struct {
	UserID         uint
	Nickname       string
	WithHardDelete bool
}
