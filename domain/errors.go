package domain

import (
	"net/http"

	"github.com/roka-crew/pkg/errors"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists").
				SetTitle(http.StatusText(http.StatusConflict)).
				SetDetail("user already exists").
				SetStatus(http.StatusConflict)
	ErrUserNotFound = errors.New("user not found").
			SetTitle(http.StatusText(http.StatusNotFound)).
			SetDetail("user not found").
			SetStatus(http.StatusNotFound)

	ErrInvalidTokenRequired = errors.New("Invalid token: Bearer token is required").
				SetTitle(http.StatusText(http.StatusUnauthorized)).
				SetDetail("Authorization header must have a Bearer token").
				SetStatus(http.StatusUnauthorized)
	ErrAuthorizationHeaderNeeded = errors.New("Authorization header is required").
					SetTitle(http.StatusText(http.StatusUnauthorized)).
					SetDetail("Missing Authorization header").
					SetStatus(http.StatusUnauthorized)
	ErrInvalidToken = errors.New("Invalid token").
			SetTitle(http.StatusText(http.StatusUnauthorized)).
			SetDetail("Invalid token").
			SetStatus(http.StatusUnauthorized)
)
