package domain

import (
	"gorm.io/gorm"
)

type Users []User

func (u Users) IsEmpty() bool {
	return len(u) == 0
}

func (u Users) First() *User {
	if u.IsEmpty() {
		return nil
	}
	return &u[0]
}

type User struct {
	gorm.Model
	Nickname   string `gorm:"uniqueIndex"`
	Resolution *string

	Gropus []Group `gorm:"many2many:user_group_mapper;"`
	Topics []Topic
}
