package store

import (
	"context"

	"github.com/roka-crew/domain"
	"github.com/roka-crew/pkg/persistence/sqlite"
	"github.com/roka-crew/presenter"
	"github.com/samber/lo"
)

type UserStore struct {
	db *sqlite.SQLite
}

func NewUserStore(
	db *sqlite.SQLite,
) *UserStore {
	return &UserStore{db: db}
}

func (s UserStore) CreateUser(ctx context.Context, params presenter.CreateUserParams) (*domain.User, error) {
	db := s.db.WithContext(ctx)

	if err := db.Create(&params).Error; err != nil {
		return nil, err
	}

	return &params, nil
}

func (s UserStore) ListUsers(ctx context.Context, params presenter.ListUsersParams) (domain.Users, error) {
	db := s.db.WithContext(ctx)

	if len(params.IDs) > 0 {
		db = db.Where("id IN ?", params.IDs)
	}

	if len(params.Nicknames) > 0 {
		db = db.Where("nickname IN ?", params.Nicknames)
	}

	if params.WithGroups {
		db = db.Preload("Groups")
	}

	if params.WithTopics {
		db = db.Preload("Topics")
	}

	if params.Limit > 0 {
		db = db.Limit(params.Limit)
	}

	var users []domain.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (s UserStore) PatchUser(ctx context.Context, params presenter.PatchUserParams) error {
	db := s.db.WithContext(ctx)

	var user domain.User
	if params.Nickname != nil {
		user.Nickname = lo.FromPtr(params.Nickname)
	}

	if params.Resolution != nil {
		user.Resolution = params.Resolution
	}

	if err := db.Updates(user).Error; err != nil {
		return err
	}

	return nil
}

func (s UserStore) DeleteUser(ctx context.Context, params presenter.DeleteUserParams) error {
	db := s.db.WithContext(ctx)

	if params.UserID > 0 {
		db = db.Where("id = ?", params.UserID)
	}

	if params.Nickname != "" {
		db = db.Where("nickname = ?", params.Nickname)
	}

	if params.WithHardDelete {
		db = db.Unscoped()
	}

	if err := db.Delete(&domain.User{}).Error; err != nil {
		return err
	}

	return nil
}
