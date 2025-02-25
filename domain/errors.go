package domain

import (
	"github.com/roka-crew/pkg/apperr"
)

var (
	// 사용자가 이미 존재하는 경우
	ErrUserAlreadyExists = apperr.New("user already exists").
				SetDetail("이미 존재하는 사용자입니다.")

	// 사용자가 존재하지 않는 경우
	ErrUserNotFound = apperr.New("user not found").
			SetDetail("사용자를 찾을 수 없습니다.")

	// 사용자가 속한 그룹이 존재하지 않는 경우
	ErrGroupNotFound = apperr.New("group not found").
				SetDetail("그룹을 찾을 수 없습니다.")

	// 구룹에 속해 있지 않은 경우
	ErrGroupNotMember = apperr.New("group not member").
				SetDetail("그룹에 속해 있지 않습니다.")

	// 미래의 목표가 이미 존재하는 경우
	ErrFutureGoalExists = apperr.New("future goal exists").
				SetDetail("미래의 목표가 이미 존재합니다.")

	// 구룹에 사용자가 있는 경우
	ErrGroupHasUsers = apperr.New("group has users").
				SetDetail("그룹에 사용자가 있습니다.")

	// 그룹에 사용자가 이미 속해있는 경우
	ErrGroupAlreadyJoined = apperr.New("group already joined").
				SetDetail("그룹에 이미 속해 있습니다.")

	// 올바르지 않은 데드라인일 경우
	ErrInvalidDeadline = apperr.New("invalid deadline").
				SetDetail("올바르지 않은 데드라인입니다.")

	//  목표가 존재하지 않는 경우
	ErrGoalNotFound = apperr.New("goal not found").
			SetDetail("목표를 찾을 수 없습니다.")

	ErrInvalidTokenRequired      = apperr.New("Invalid token: Bearer token is required")
	ErrAuthorizationHeaderNeeded = apperr.New("Authorization header is required")
	ErrInvalidToken              = apperr.New("Invalid token")
)
