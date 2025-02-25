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

func (s UserService) CreateUser(ctx context.Context, request presenter.CreateUserRequest) (presenter.CreateUserResponse, error) {
	// 1. 이름 중복 확인
	// listUsers, err := s.userStore.ListUsers(ctx, presenter.ListUsersParams{
	// 	Nicknames: []string{request.Nickname},
	// 	Limit:     1,
	// })
	// if err != nil {
	// 	return presenter.CreateUserResponse{}, err
	// }

	// if !listUsers.IsEmpty() {
	// 	return presenter.CreateUserResponse{}, domain.ErrUserAlreadyExists
	// }

	createdUser, err := s.userStore.CreateUser(ctx, presenter.CreateUserParams{
		Nickname:   request.Nickname,
		Resolution: request.Resolution,
	})
	if err != nil {
		return presenter.CreateUserResponse{}, err
	}

	return presenter.CreateUserResponse{
		Nickname:   createdUser.Nickname,
		Resolution: createdUser.Resolution,
	}, nil
}

func (s UserService) FindUserByMe(ctx context.Context, request presenter.FindUserByMeRequest) (presenter.FindUserByMeResponse, error) {
	listUsers, err := s.userStore.ListUsers(ctx, presenter.ListUsersParams{
		IDs:   []uint{request.RequestUserID},
		Limit: 1,
	})
	if err != nil {
		return presenter.FindUserByMeResponse{}, err
	}
	if listUsers.IsEmpty() {
		return presenter.FindUserByMeResponse{}, domain.ErrUserNotFound
	}

	return presenter.FindUserByMeResponse{
		Nickname:   listUsers.First().Nickname,
		Resolution: listUsers.First().Resolution,
	}, nil
}

func (s UserService) PatchUserByMe(ctx context.Context, request presenter.PatchUserByMeRequest) error {
	// 1. 해당 사용자가 존재하는지 확인
	listUsers, err := s.userStore.ListUsers(ctx, presenter.ListUsersParams{
		IDs:   []uint{request.RequestUserID},
		Limit: 1,
	})
	if err != nil {
		return err
	}
	if listUsers.IsEmpty() {
		return domain.ErrUserNotFound
	}

	// 2. 사용자 정보 수정
	err = s.userStore.PatchUser(ctx, presenter.PatchUserParams{
		UserID:     &request.RequestUserID,
		Nickname:   request.Nickname,
		Resolution: request.Resolution,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s UserService) DeleteUserByMe(ctx context.Context, request presenter.DeleteUserByMeRequest) error {
	// 1. 해당 사용자가 존재하는지 확인
	listUsers, err := s.userStore.ListUsers(ctx, presenter.ListUsersParams{
		IDs:   []uint{request.RequestUserID},
		Limit: 1,
	})
	if err != nil {
		return err
	}
	if listUsers.IsEmpty() {
		return domain.ErrUserNotFound
	}

	// 2. 사용자 정보 삭제
	err = s.userStore.DeleteUser(ctx, presenter.DeleteUserParams{
		UserID: request.RequestUserID,
	})
	if err != nil {
		return err
	}

	return nil
}
