package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/roka-crew/domain"
	"github.com/roka-crew/internal/user/service"
	"github.com/roka-crew/pkg/ctxutil"
	"github.com/roka-crew/presenter"
	"github.com/roka-crew/router"
	"github.com/roka-crew/router/middleware"
)

type UserHandler struct {
	userService    *service.UserService
	router         *router.Router
	authMiddleware *middleware.AuthMiddleware
	ctxutil        *ctxutil.CtxUtil
}

func NewUserHandler(
	userService *service.UserService,
	router *router.Router,
	authMiddleware *middleware.AuthMiddleware,
	ctxutil *ctxutil.CtxUtil,
) *UserHandler {
	userHandler := &UserHandler{
		userService:    userService,
		router:         router,
		authMiddleware: authMiddleware,
		ctxutil:        ctxutil,
	}

	users := router.Group("/users")
	{
		users.POST("/", userHandler.CreateUser)
		users.GET("/me", userHandler.FindUserByMe, authMiddleware.AuthenticateRequest)
		// users.PATCH("/", userHandler.CreateUser)
		// users.DELETE("/", userHandler.CreateUser)
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

func (h UserHandler) FindUserByMe(c echo.Context) error {
	var (
		request presenter.FindUserByMeRequest
		err     error
	)

	request.RequestUserID, err = h.ctxutil.GetRequestUserID(c)
	if err != nil {
		return err
	}

	foundUserByMe, err := h.userService.FindUserByMe(c.Request().Context(), request)
	if err != nil {
		return err
	}

	switch {
	case err == nil:
		return c.JSON(http.StatusOK, presenter.NewFindUserByMeRequest(foundUserByMe))
	default:
		return err
	}
}
