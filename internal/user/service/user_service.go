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

func (s UserService) CreateUser(ctx context.Context, request presenter.CreateUserRequest) (*domain.User, error) {
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

	return nil, nil
}
