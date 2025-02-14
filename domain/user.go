package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname   string `gorm:"uniqueIndex"`
	Resolution *string

	Gropus []Group `gorm:"many2many:user_group_mapper;"`
	Topics []Topic
}

type CreateUserParams = User

type ListUsersParams struct {
	IDs        []int
	Nicknames  []string
	WithGroups bool
	WithTopics bool
}

type PatchUserParams struct {
	UserID     *int
	Nickname   string
	Resolution *string
}

type DeleteUserParams struct {
	UserID         int
	Nickname       string
	WithHardDelete bool
}
