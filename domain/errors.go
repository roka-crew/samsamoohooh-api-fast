package domain

import (
	"github.com/roka-crew/pkg/errors"
)

var (
	// 사용자가 이미 존재하는 경우
	ErrUserAlreadyExists = errors.New("user already exists")

	// 사용자가 존재하지 않는 경우
	ErrUserNotFound = errors.New("user not found")

	// 사용자가 속한 그룹이 존재하지 않는 경우
	ErrGroupNotFound = errors.New("group not found")

	// 구룹에 속해 있지 않은 경우
	ErrGroupNotMember = errors.New("group not member")

	// 미래의 목표가 이미 존재하는 경우
	ErrFutureGoalExists = errors.New("future goal exists")

	// 구룹에 사용자가 있는 경우
	ErrGroupHasUsers = errors.New("group has users")

	// 그룹에 사용자가 이미 속해있는 경우
	ErrGroupAlreadyJoined = errors.New("group already joined")

	ErrInvalidTokenRequired      = errors.New("Invalid token: Bearer token is required")
	ErrAuthorizationHeaderNeeded = errors.New("Authorization header is required")
	ErrInvalidToken              = errors.New("Invalid token")
)
