package service

import (
	"context"

	"github.com/roka-crew/internal/group/store"
	"github.com/roka-crew/presenter"
)

type GroupService struct {
	groupStore *store.GroupStore
}

func NewGroupService(groupStore *store.GroupStore) *GroupService {
	return &GroupService{
		groupStore: groupStore,
	}
}

func (s GroupService) CreateGroup(ctx context.Context, request presenter.CreateGroupRequest) (*presenter.CreateGroupResponse, error) {
	// 1. 모임 생성
	createdGroup, err := s.groupStore.CreateGroup(ctx, presenter.CreateGroupParams{
		BookTitle:        request.BookTitle,
		BookAuthor:       request.BookAuthor,
		BookPageMax:      request.BookPageMax,
		BookPageCount:    request.BookPageCount,
		BookPublisher:    request.BookPublisher,
		BookIntroduction: request.BookIntroduction,
	})
	if err != nil {
		return nil, err
	}

	// 2. 이미 추가된 사용자인지 확인
	s.groupStore.ListGroupUsers(ctx, presenter.ListGroupUsersParams{})

	// 2. n : m 관계 추가
	err = s.groupStore.AddGroupUsers(ctx, presenter.AddGroupUsersParams{
		GroupID: createdGroup.ID,
		UserIDs: []uint{request.RequestUserID},
	})
	if err != nil {
		return nil, err
	}

	return &presenter.CreateGroupResponse{
		BookTitle:        createdGroup.BookTitle,
		BookAuthor:       createdGroup.BookAuthor,
		BookPageMax:      createdGroup.BookPageMax,
		BookPageCount:    createdGroup.BookPageCount,
		BookPublisher:    createdGroup.BookPublisher,
		BookIntroduction: createdGroup.BookIntroduction,
	}, nil
}
