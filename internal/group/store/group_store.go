package store

import (
	"context"

	"github.com/roka-crew/domain"
	"github.com/roka-crew/pkg/persistence/sqlite"
	"github.com/roka-crew/presenter"
	"gorm.io/gorm"
)

type GroupStore struct {
	db *sqlite.SQLite
}

func NewGormStore(db *sqlite.SQLite) *GroupStore {
	return &GroupStore{
		db: db,
	}
}

func (s GroupStore) CreateGroup(ctx context.Context, params presenter.CreateGroupParams) (*domain.Group, error) {
	db := s.db.WithContext(ctx)

	if err := db.Create(&params).Error; err != nil {
		return nil, err
	}

	return &params, nil
}

func (s GroupStore) ListGroups(ctx context.Context, params presenter.ListGroupsParams) (domain.Groups, error) {
	db := s.db.WithContext(ctx)

	if len(params.IDs) > 0 {
		db = db.Where("id IN ?", params.IDs)
	}

	if params.WithUsers {
		db = db.Preload("Users")
	}

	if params.WithGoals {
		db = db.Preload("Goals")
	}

	if params.Limit > 0 {
		db = db.Limit(params.Limit)
	}

	var groups domain.Groups
	if err := db.Find(&groups).Error; err != nil {
		return nil, err
	}

	return groups, nil
}

func (s GroupStore) PatchGroup(ctx context.Context) error {
	return nil
}

func (s GroupStore) DeleteGroup(ctx context.Context) error {
	return nil
}

func (s GroupStore) ListGroupUsers(ctx context.Context, params presenter.ListGroupUsersParams) (domain.Users, error) {
	db := s.db.WithContext(ctx)

	var users domain.Users
	if err := db.Model(&domain.Group{Model: gorm.Model{ID: params.GroupID}}).Association("Users").Find(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func (s GroupStore) AddGroupUsers(ctx context.Context, params presenter.AddGroupUsersParams) error {
	db := s.db.WithContext(ctx)

	// TODO: 어떻게 처리할 지 고민하기
	// 1. users 조회
	var users domain.Users
	if err := db.Where("id IN ?", params.UserIDs).Find(&users).Error; err != nil {
		return err
	}

	if err := db.Model(&domain.Group{Model: gorm.Model{ID: params.GroupID}}).Association("Users").Append(&users); err != nil {
		return err
	}

	return nil
}

func (s GroupStore) RemoveUser(ctx context.Context) error {
	return nil
}
