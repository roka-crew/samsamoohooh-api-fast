package store

import (
	"context"
	"fmt"

	"github.com/roka-crew/domain"
	"github.com/roka-crew/pkg/errors"
	"github.com/roka-crew/pkg/persistence/sqlite"
	"github.com/roka-crew/presenter"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type GoalStore struct {
	db *sqlite.SQLite
}

func NewGoalStore(db *sqlite.SQLite) *GoalStore {
	return &GoalStore{
		db: db,
	}
}

func (s GoalStore) CreateGoal(ctx context.Context, params presenter.CreateGoalParams) (domain.Goal, error) {
	db := s.db.WithContext(ctx)

	if err := db.Create(&params).Error; err != nil {
		return domain.Goal{}, errors.InteralError(err)
	}

	return params, nil
}

func (s GoalStore) ListGoals(ctx context.Context, params presenter.ListGoalsParams) (domain.Goals, error) {
	db := s.db.WithContext(ctx)

	if len(params.IDs) > 0 {
		db = db.Where("id IN ?", params.IDs)
	}

	if len(params.GroupIDs) > 0 {
		db = db.Where("group_id IN ?", params.GroupIDs)
	}

	if params.DeadlineOrder != "" {
		db = db.Order(fmt.Sprintf("deadline %s", params.DeadlineOrder))
	}

	if params.WithTopics {
		db = db.Preload("Topics")
	}

	if params.Limit > 0 {
		db = db.Limit(params.Limit)
	}

	if !params.GteDeadline.IsZero() {
		db = db.Where("deadline >= ?", params.GteDeadline)
	}

	var goals domain.Goals
	if err := db.Find(&goals).Error; err != nil {
		return domain.Goals{}, errors.InteralError(err)
	}

	return goals, nil
}

func (s GoalStore) PatchGoal(ctx context.Context, params presenter.PatchGoalParams) error {
	db := s.db.WithContext(ctx)

	var updates domain.Goal

	if params.Deadline != nil {
		updates.Deadline = lo.FromPtr(params.Deadline)
	}

	if params.PageRange != nil {
		updates.PageRange = lo.FromPtr(params.PageRange)
	}

	if err := db.Where("id = ?", params.ID).Updates(&updates).Error; err != nil {
		return errors.InteralError(err)
	}

	return nil
}

func (s GoalStore) DeleteGoal(ctx context.Context, params presenter.DeleteGoalParams) error {
	db := s.db.WithContext(ctx)

	if err := db.Delete(&domain.Goal{Model: gorm.Model{ID: params.GoalID}}, params.GoalID).Error; err != nil {
		return errors.InteralError(err)
	}

	return nil
}
