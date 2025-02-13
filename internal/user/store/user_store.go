package store

import (
	"github.com/roka-crew/pkg/persistence"
)

type UserStore struct {
	db *persistence.Storm
}

func NewUserStore(
	db *persistence.Storm,
) *UserStore {
	return &UserStore{db: db}
}
