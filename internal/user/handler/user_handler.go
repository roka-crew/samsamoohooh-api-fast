package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/roka-crew/domain"
	"github.com/roka-crew/internal/user/service"
	"github.com/roka-crew/presenter"
	"github.com/roka-crew/router"
)

type UserHandler struct {
	userService *service.UserService
	router      *router.Router
}

func NewUserHandler(
	userService *service.UserService,
	router *router.Router,
) *UserHandler {
	userHandler := &UserHandler{
		userService: userService,
		router:      router,
	}

	users := router.Group("/users")
	{
		users.POST("/", userHandler.CreateUser)
	}

	return userHandler
}

func (h UserHandler) CreateUser(c echo.Context) error {
	var (
		request presenter.CreateUserRequest
		err     error
	)

	if err = c.Bind(&request); err != nil {
		return err
	}

	createdUser, err := h.userService.CreateUser(c.Request().Context(), request)

	switch {
	case err == nil:
		return c.JSON(http.StatusCreated, presenter.NewCreateUserResponse(createdUser))
	case errors.Is(err, domain.ErrUserAlreadyExists):
		return c.NoContent(http.StatusConflict)
	default:
		return err
	}
}
