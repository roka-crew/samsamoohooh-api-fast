package domain

import (
	"github.com/roka-crew/pkg/errors"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")

	ErrGroupNotFound      = errors.New("group not found")
	ErrGroupNotOwned      = errors.New("group not owned")
	ErrGroupHasUsers      = errors.New("group has users")
	ErrGroupAlreadyJoined = errors.New("group already joined")

	ErrInvalidTokenRequired      = errors.New("Invalid token: Bearer token is required")
	ErrAuthorizationHeaderNeeded = errors.New("Authorization header is required")
	ErrInvalidToken              = errors.New("Invalid token")
)
