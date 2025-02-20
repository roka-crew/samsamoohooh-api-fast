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

func (s GroupService) PatchGroup(ctx context.Context, request presenter.PatchGroupRequest) error {
	// 1. 구룹이 존재하는지 확인
	listGroup, err := s.groupStore.ListGroups(ctx, presenter.ListGroupsParams{
		IDs:   []uint{request.GroupID},
		Limit: 1,
	})
	if err != nil {
		return err
	}
	if listGroup.IsEmpty() {
		return domain.ErrGroupNotFound
	}

	// 2. 수정 요청한 모임이 내 모임이 맞는지 확인
	exists, err := s.groupStore.ExistsGroupUser(ctx, presenter.ExistsGroupUserParams{
		GroupID: request.GroupID,
		UserID:  request.RequestUserID,
	})
	if err != nil {
		return err
	}
	if !exists {
		return domain.ErrGroupNotMember
	}

	// 3s. 모임 수정
	err = s.groupStore.PatchGroup(ctx, presenter.PatchGroupParams{
		GroupID:          request.GroupID,
		BookTitle:        request.BookTitle,
		BookAuthor:       request.BookAuthor,
		BookPageMax:      request.BookPageMax,
		BookPageCount:    request.BookPageCount,
		BookPublisher:    request.BookPublisher,
		BookIntroduction: request.BookIntroduction,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s GroupService) DeleteGroup(ctx context.Context, reqeust presenter.DeleteGroupRequest) error {
	// 1. 구룹이 존재하는지 확인
	listGroup, err := s.groupStore.ListGroups(ctx, presenter.ListGroupsParams{
		IDs:       []uint{reqeust.GroupID},
		WithUsers: true,
		Limit:     1,
	})
	if err != nil {
		return err
	}

	if listGroup.IsEmpty() {
		return domain.ErrGroupNotFound
	}

	// 2. 삭제 요청한 모임이 내 모임이 맞는지 확인
	exists, err := s.groupStore.ExistsGroupUser(ctx, presenter.ExistsGroupUserParams{
		GroupID: reqeust.GroupID,
		UserID:  reqeust.RequestUserID,
	})
	if err != nil {
		return err
	}
	if !exists {
		return domain.ErrGroupNotMember
	}

	// 3. 해당 구룹의 사용자들이 있는지 확인
	if len(listGroup.First().Users) >= 1 {
		return domain.ErrGroupHasUsers
	}

	// 4. 해당 구룹 삭제
	err = s.groupStore.DeleteGroup(ctx, presenter.DeleteGroupParams{
		GroupID: reqeust.GroupID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s GroupService) LeaveGroup(ctx context.Context, reqeust presenter.LeaveGroupRequest) error {
	// 1. 구룹이 존재하는지 확인
	listGroup, err := s.groupStore.ListGroups(ctx, presenter.ListGroupsParams{
		IDs:       []uint{reqeust.GroupID},
		Limit:     1,
		WithUsers: true,
	})
	if err != nil {
		return err
	}

	if listGroup.IsEmpty() {
		return domain.ErrGroupNotFound
	}

	// 2. 나가려는 모임이 내 모임이 맞는지 확인
	exists, err := s.groupStore.ExistsGroupUser(ctx, presenter.ExistsGroupUserParams{
		GroupID: reqeust.GroupID,
		UserID:  reqeust.RequestUserID,
	})
	if err != nil {
		return err
	}
	if !exists {
		return domain.ErrGroupNotMember
	}

	// 3. 해당 구룹에서 나가기
	err = s.groupStore.RemoveGroupUser(ctx, presenter.RemoveGroupUserParams{
		GroupID: reqeust.GroupID,
		UserID:  reqeust.RequestUserID,
	})
	if err != nil {
		return err
	}

	// 4. 해당 구룹의 사용자들이 있는지 확인
	if len(listGroup.First().Users) < 1 {
		// 5. 해당 구룹 삭제
		err = s.groupStore.DeleteGroup(ctx, presenter.DeleteGroupParams{
			GroupID: reqeust.GroupID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s GroupService) JoinGroup(ctx context.Context, request presenter.JoinGroupRequest) error {
	// 1. 구룹이 존재하는지 확인
	listGroup, err := s.groupStore.ListGroups(ctx, presenter.ListGroupsParams{
		IDs:   []uint{request.GroupID},
		Limit: 1,
	})
	if err != nil {
		return err
	}
	if listGroup.IsEmpty() {
		return domain.ErrGroupNotFound
	}

	// 2. 이미 가입한 모임인지 확인
	exists, err := s.groupStore.ExistsGroupUser(ctx, presenter.ExistsGroupUserParams{
		GroupID: request.GroupID,
		UserID:  request.RequestUserID,
	})
	if err != nil {
		return err
	}
	if exists {
		return domain.ErrGroupAlreadyJoined
	}

	// 3. 모임 가입
	err = s.groupStore.AddGroupUsers(ctx, presenter.AddGroupUsersParams{
		GroupID: request.GroupID,
		UserIDs: []uint{request.RequestUserID},
	})
	if err != nil {
		return err
	}

	return nil
}
