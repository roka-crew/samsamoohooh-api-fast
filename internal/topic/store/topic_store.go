package store

import (
	"context"

	"github.com/roka-crew/domain"
	"github.com/roka-crew/pkg/errors"
	"github.com/roka-crew/pkg/persistence/sqlite"
	"github.com/roka-crew/presenter"
	"github.com/samber/lo"
)

type TopicStore struct {
	db *sqlite.SQLite
}

func NewTopicStore(db *sqlite.SQLite) *TopicStore {
	return &TopicStore{
		db: db,
	}
}

func (s TopicStore) CreateTopic(ctx context.Context, params presenter.CreateTopicParams) (domain.Topic, error) {
	db := s.db.WithContext(ctx)

	if err := db.Create(&params).Error; err != nil {
		return domain.Topic{}, errors.InteralError(err)
	}

	return domain.Topic{}, nil
}

func (s TopicStore) ListTopics(ctx context.Context, params presenter.ListTopicsParams) (domain.Topics, error) {
	db := s.db.WithContext(ctx)

	if len(params.IDs) > 0 {
		db = db.Where("id IN ?", params.IDs)
	}

	if len(params.GoalIDs) > 0 {
		db = db.Where("goal_id IN ?", params.GoalIDs)
	}

	if len(params.UserIDs) > 0 {
		db = db.Where("user_id IN ?", params.UserIDs)
	}

	if params.Limit > 0 {
		db = db.Limit(params.Limit)
	}

	var topics domain.Topics
	if err := db.Find(&topics).Error; err != nil {
		return domain.Topics{}, errors.InteralError(err)
	}

	return topics, nil
}

func (s TopicStore) PatchTopic(ctx context.Context, params presenter.PatchTopicParams) error {
	db := s.db.WithContext(ctx)

	var updates = domain.Topic{}

	if params.Content != nil {
		updates.Content = lo.FromPtr(params.Content)
	}

	if params.Title != nil {
		updates.Title = lo.FromPtr(params.Title)
	}

	if err := db.Model(&domain.Topic{}).Where("id = ?", params.ID).Updates(updates).Error; err != nil {
		return errors.InteralError(err)
	}

	return nil
}

func (s TopicStore) DeleteTopic(ctx context.Context, params presenter.DeleteTopicParams) error {
	db := s.db.WithContext(ctx)

	if err := db.Where("id = ?", params.ID).Delete(&domain.Topic{}).Error; err != nil {
		return errors.InteralError(err)
	}

	return nil
}
