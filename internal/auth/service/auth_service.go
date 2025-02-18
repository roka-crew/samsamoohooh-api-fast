package service

import (
	"context"

	"github.com/roka-crew/domain"
	"github.com/roka-crew/internal/user/store"
	"github.com/roka-crew/pkg/token"
	"github.com/roka-crew/presenter"
	"github.com/roka-crew/router/middleware"
)

type AuthService struct {
	tokenService *token.Token
	userStore    *store.UserStore
}

func NewAuthService(
	tokenService *token.Token,
	userStore *store.UserStore,
	authMiddleware *middleware.AuthMiddleware,
) *AuthService {
	return &AuthService{
		tokenService: tokenService,
		userStore:    userStore,
	}
}

func (s AuthService) IssueToken(ctx context.Context, request presenter.IssueTokenRequest) (*presenter.IssueTokenResponse, error) {
	// 1. 닉네임으로 사용자를 조회한다.
	listUsers, err := s.userStore.ListUsers(ctx, presenter.ListUsersParams{
		Nicknames: []string{request.Nickname},
		Limit:     1,
	})
	if err != nil {
		return nil, err
	}
	if listUsers.IsEmpty() {
		return nil, domain.ErrUserNotFound
	}

	// 2. 토큰을 발급한다.
	tokenString, err := s.tokenService.GenerateToken(listUsers.First().ID)
	if err != nil {
		return nil, err
	}

	return &presenter.IssueTokenResponse{
		Token: tokenString,
	}, nil
}

func (s AuthService) Validate(ctx context.Context, request presenter.ValidateRequest) (*presenter.ValidateResponse, error) {
	// 1. 사용자를 조회한다.
	listUsers, err := s.userStore.ListUsers(ctx, presenter.ListUsersParams{
		Limit: 1,
		IDs:   []uint{request.RequestUserID},
	})
	if err != nil {
		return nil, err
	}
	if listUsers.IsEmpty() {
		return nil, domain.ErrUserNotFound
	}

	// 2. 사용자의 닉네임을 반환한다.
	return &presenter.ValidateResponse{
		Nickname:   listUsers.First().Nickname,
		Resolution: listUsers.First().Resolution,
	}, nil
}
