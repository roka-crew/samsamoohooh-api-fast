package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/roka-crew/domain"
	"github.com/roka-crew/internal/user/service"
	"github.com/roka-crew/pkg/ctxutil"
	"github.com/roka-crew/pkg/errors"
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
		users.PATCH("/me", userHandler.PatchUserByMe, authMiddleware.AuthenticateRequest)
		users.DELETE("/me", userHandler.DeleteUserByMe, authMiddleware.AuthenticateRequest)
	}

	return userHandler
}

// CreateUser
// @Summary CreateUser - 사용자 단건 생성✅
// @Tags users
// @Param CreateUserRequest body presenter.CreateUserRequest true "사용자 생성 요청"
// @Produce json
// @Success 201 {object} presenter.CreateUserResponse"회원가입에 성공한 사용자 정보"
// @Failure 409 {object} errors.Error "사용자가 이미 존재함 : user already exists"
// @Router /users [post]
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
		return errors.Restore(err).SetStatus(http.StatusConflict)
	default:
		return err
	}
}

// FindUserByMe
// @Summary FindUserByMe - 사용자(나) 단건 조회 ✅
// @Tags users
// @Produce json
// @Success 200 {object} presenter.FindUserByMeResponse "내 정보 조회에 성공한 사용자 정보"
// @Failure 404 {object} errors.Error "유저를 찾지 못함: user not found"
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
		return errors.Restore(err).SetStatus(http.StatusNotFound)
	default:
		return err
	}

}

// PatchUserByMe
// @Summary PatchUserByMe - 사용자(나) 단건 수정✅
// @Tags users
// @Produce json
// @Param PatchUserByMeRequest body presenter.PatchUserByMeRequest true "사용자 수정 요청"
// @Success 204
// @Failure 404 {object} errors.Error "유저를 찾지 못함: user not found"
// @Router /users/me [get]
// @Security BearerAuth
func (h UserHandler) PatchUserByMe(c echo.Context) error {
	var (
		request presenter.PatchUserByMeRequest
		err     error
	)

	if err = c.Bind(&request); err != nil {
		return err
	}

	request.RequestUserID, err = h.ctxutil.GetRequestUserID(c)
	if err != nil {
		return err
	}

	err = h.userService.PatchUserByMe(c.Request().Context(), request)

	switch {
	case err == nil:
		return c.NoContent(http.StatusNoContent)
	case errors.Is(err, domain.ErrUserNotFound):
		return errors.Restore(err).SetStatus(http.StatusNotFound)
	default:
		return err
	}
}

// DeleteUserByMe
// @Summary DeleteUserByMe - 사용자(나) 단건 삭제✅
// @Tags users
// @Produce json
// @Success 204
// @Failure 404 {object} errors.Error "유저를 찾지 못함: user not found"
// @Router /users/me [delete]
// @Security BearerAuth
func (h UserHandler) DeleteUserByMe(c echo.Context) error {
	var (
		reqeust presenter.DeleteUserByMeRequest
		err     error
	)

	reqeust.RequestUserID, err = h.ctxutil.GetRequestUserID(c)
	if err != nil {
		return err
	}

	err = h.userService.DeleteUserByMe(c.Request().Context(), reqeust)

	switch {
	case err == nil:
		return c.NoContent(http.StatusNoContent)
	case errors.Is(err, domain.ErrUserNotFound):
		return errors.Restore(err).SetStatus(http.StatusNotFound)
	default:
		return err
	}
}
