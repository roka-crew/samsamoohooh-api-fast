package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/roka-crew/internal/user/presenter"
	"github.com/roka-crew/internal/user/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(
	userService *service.UserService,
) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h UserHandler) CreateUser(c echo.Context) error {
	var (
		req presenter.CreateUserRequest
		_   presenter.CreateUserResponse
		err error
	)

	if err = c.Bind(&req); err != nil {
		return err
	}

	return nil
}
