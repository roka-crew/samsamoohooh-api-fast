package service

import (
	"context"

	"github.com/roka-crew/domain"
	groupStore "github.com/roka-crew/internal/group/store"
	userStore "github.com/roka-crew/internal/user/store"
	"github.com/roka-crew/presenter"
)

type GroupService struct {
	groupStore *groupStore.GroupStore
	userStore  *userStore.UserStore
}

func NewGroupService(
	groupStore *groupStore.GroupStore,
	userStore *userStore.UserStore,
) *GroupService {
	return &GroupService{
		groupStore: groupStore,
		userStore:  userStore,
	}
}

func (s GroupService) CreateGroup(ctx context.Context, request presenter.CreateGroupRequest) (presenter.CreateGroupResponse, error) {
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
		return presenter.CreateGroupResponse{}, err
	}

	// 2. n : m 관계 추가
	err = s.groupStore.AddGroupUsers(ctx, presenter.AddGroupUsersParams{
		GroupID: createdGroup.ID,
		UserIDs: []uint{request.RequestUserID},
	})
	if err != nil {
		return presenter.CreateGroupResponse{}, err
	}

	return presenter.CreateGroupResponse{
		BookTitle:        createdGroup.BookTitle,
		BookAuthor:       createdGroup.BookAuthor,
		BookPageMax:      createdGroup.BookPageMax,
		BookPageCount:    createdGroup.BookPageCount,
		BookPublisher:    createdGroup.BookPublisher,
		BookIntroduction: createdGroup.BookIntroduction,
	}, nil
}

func (s GroupService) ListGroups(ctx context.Context, request presenter.ListGroupsRequest) (presenter.ListGroupsResponse, error) {
	listGroups, err := s.userStore.ListUserGroups(ctx, presenter.ListUserGroupsParams{
		UserID: request.RequestUserID,
	})
	if err != nil {
		return presenter.ListGroupsResponse{}, err
	}

	if listGroups.IsEmpty() {
		return presenter.ListGroupsResponse{}, domain.ErrGroupNotFound
	}

	responseGroups := make([]presenter.GroupsResponse, 0, len(listGroups))
	for _, group := range listGroups {
		responseGroups = append(responseGroups, presenter.GroupsResponse{
			ID:               group.ID,
			BookTitle:        group.BookTitle,
			BookAuthor:       group.BookAuthor,
			BookPageMax:      group.BookPageMax,
			BookPageCount:    group.BookPageCount,
			BookPublisher:    group.BookPublisher,
			BookIntroduction: group.BookIntroduction,
		})
	}

	return presenter.ListGroupsResponse{
		Groups: responseGroups,
	}, nil
}
