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

// CreateUser
// @Tags users
// @Summary CreateUser ✅
// @Produce json
// @Router /users [post]
// @Success 201 {object} presenter.CreateUserRequest "회원가입에 성공한 사용자 정보"
// @Security BearerAuth
func (h UserHandler) CreateUser(c echo.Context) error {
	var (
		request  presenter.CreateUserRequest
		response *presenter.CreateUserResponse
		err      error
	)

	if err = c.Bind(&request); err != nil {
		return err
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	response, err = h.userService.CreateUser(c.Request().Context(), request)

	switch {
	case err == nil:
		return c.JSON(http.StatusCreated, response)
	case errors.Is(err, domain.ErrUserAlreadyExists):
		return c.NoContent(http.StatusConflict)
	default:
		return err
	}
}

// FindUserByMe
// @Tags users
// @Summary FindUserByMe ✅
// @Produce json
// @Success 200 {object} presenter.FindUserByMeResponse "내 정보 조회에 성공한 사용자 정보"
// @Router /users/me [get]
// @Security BearerAuth
func (h UserHandler) FindUserByMe(c echo.Context) error {
	var (
		request  presenter.FindUserByMeRequest
		response *presenter.FindUserByMeResponse
		err      error
	)

	request.RequestUserID, err = h.ctxutil.GetRequestUserID(c)
	if err != nil {
		return err
	}

	response, err = h.userService.FindUserByMe(c.Request().Context(), request)

	switch {
	case err == nil:
		return c.JSON(http.StatusOK, response)
	case errors.Is(err, domain.ErrUserNotFound):
		return c.NoContent(http.StatusNotFound)
	default:
		return err
	}
}
