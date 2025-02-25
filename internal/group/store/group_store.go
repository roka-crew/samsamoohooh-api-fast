package store

import (
	"context"

	"github.com/roka-crew/domain"
	"github.com/roka-crew/pkg/apperr"
	"github.com/roka-crew/pkg/persistence/sqlite"
	"github.com/roka-crew/presenter"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type GroupStore struct {
	db *sqlite.SQLite
}

func NewGroupStore(db *sqlite.SQLite) *GroupStore {
	return &GroupStore{
		db: db,
	}
}

func (s GroupStore) CreateGroup(ctx context.Context, params presenter.CreateGroupParams) (domain.Group, error) {
	db := s.db.WithContext(ctx)

	if err := db.Create(&params).Error; err != nil {
		return domain.Group{}, apperr.NewInternalError(err)
	}

	return params, nil
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
		return nil, apperr.NewInternalError(err)
	}

	return groups, nil
}

func (s GroupStore) PatchGroup(ctx context.Context, params presenter.PatchGroupParams) error {
	db := s.db.WithContext(ctx)

	var updates domain.Group

	if params.BookTitle != nil {
		updates.BookTitle = lo.FromPtr(params.BookTitle)
	}

	if params.BookAuthor != nil {
		updates.BookAuthor = lo.FromPtr(params.BookAuthor)
	}

	if params.BookPageMax != nil {
		updates.BookPageMax = lo.FromPtr(params.BookPageMax)
	}

	if params.BookPageCount != nil {
		updates.BookPageCount = lo.FromPtr(params.BookPageCount)
	}

	if params.BookPublisher != nil {
		updates.BookPublisher = params.BookPublisher
	}

	if params.BookIntroduction != nil {
		updates.BookIntroduction = params.BookIntroduction
	}

	if err := db.Model(&domain.Group{}).Where("id = ?", params.GroupID).Updates(&updates).Error; err != nil {
		return apperr.NewInternalError(err)
	}

	return nil
}

func (s GroupStore) DeleteGroup(ctx context.Context, params presenter.DeleteGroupParams) error {
	db := s.db.WithContext(ctx)

	if err := db.Delete(&domain.Group{Model: gorm.Model{ID: params.GroupID}}).Error; err != nil {
		return apperr.NewInternalError(err)
	}

	return nil
}

func (s GroupStore) AddGroupUsers(ctx context.Context, params presenter.AddGroupUsersParams) error {
	db := s.db.WithContext(ctx)

	// 1. users 조회
	var users domain.Users
	if err := db.Where("id IN ?", params.UserIDs).Find(&users).Error; err != nil {
		return apperr.NewInternalError(err)
	}

	if err := db.Model(&domain.Group{Model: gorm.Model{ID: params.GroupID}}).Association("Users").Append(&users); err != nil {
		return apperr.NewInternalError(err)
	}

	return nil
}

func (s GroupStore) RemoveGroupUser(ctx context.Context, params presenter.RemoveGroupUserParams) error {
	db := s.db.WithContext(ctx)

	if err := db.Model(&domain.Group{Model: gorm.Model{ID: params.GroupID}}).Association("Users").Delete(&domain.User{Model: gorm.Model{ID: params.UserID}}); err != nil {
		return apperr.NewInternalError(err)
	}

	return nil
}

func (s GroupStore) ExistsGroupUser(ctx context.Context, params presenter.ExistsGroupUserParams) (bool, error) {
	db := s.db.WithContext(ctx)

	var count int64
	if err := db.Table("user_group_mapper").
		Where("user_id = ?", params.UserID).
		Where("group_id = ?", params.GroupID).
		Limit(1).
		Count(&count).Error; err != nil {
		return false, apperr.NewInternalError(err)
	}

	return count > 0, nil
}
