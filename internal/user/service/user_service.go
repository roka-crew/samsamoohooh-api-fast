package service

import (
	"context"

	"github.com/roka-crew/domain"
	"github.com/roka-crew/internal/user/store"
	"github.com/roka-crew/presenter"
)

type UserService struct {
	userStore *store.UserStore
}

func NewUserService(userStore *store.UserStore) *UserService {
	return &UserService{userStore: userStore}
}

func (s UserService) CreateUser(ctx context.Context, request presenter.CreateUserRequest) (*presenter.CreateUserResponse, error) {
	// 1. 이름 중복 확인
	listUsers, err := s.userStore.ListUsers(ctx, presenter.ListUsersParams{
		Nicknames: []string{request.Nickname},
		Limit:     1,
	})
	if err != nil {
		return nil, err
	}

	if !listUsers.IsEmpty() {
		return nil, domain.ErrUserAlreadyExists
	}

	createdUser, err := s.userStore.CreateUser(ctx, presenter.CreateUserParams{
		Nickname:   request.Nickname,
		Resolution: request.Resolution,
	})
	if err != nil {
		return nil, err
	}

	return &presenter.CreateUserResponse{
		Nickname:   createdUser.Nickname,
		Resolution: createdUser.Resolution,
	}, nil
}

func (u UserService) FindUserByMe(ctx context.Context, reqeust presenter.FindUserByMeRequest) (*presenter.FindUserByMeResponse, error) {
	listUsers, err := u.userStore.ListUsers(ctx, presenter.ListUsersParams{
		IDs:   []uint{reqeust.RequestUserID},
		Limit: 1,
	})
	if err != nil {
		return nil, err
	}
	if listUsers.IsEmpty() {
		return nil, domain.ErrUserNotFound
	}

	return &presenter.FindUserByMeResponse{
		Nickname:   listUsers.First().Nickname,
		Resolution: listUsers.First().Resolution,
	}, nil
}
