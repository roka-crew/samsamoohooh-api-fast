package handler

import (
	"net/http"

	"github.com/roka-crew/pkg/ctxutil"
	"github.com/roka-crew/pkg/errors"
	"github.com/roka-crew/router/middleware"

	"github.com/labstack/echo/v4"
	"github.com/roka-crew/domain"
	"github.com/roka-crew/internal/auth/service"
	"github.com/roka-crew/presenter"
	"github.com/roka-crew/router"
)

type AuthHandler struct {
	router         *router.Router
	authService    *service.AuthService
	authMiddleware *middleware.AuthMiddleware
	ctxutil        *ctxutil.CtxUtil
}

func NewAuthHandler(
	router *router.Router,
	authService *service.AuthService,
	authMiddleware *middleware.AuthMiddleware,
	ctxUtil *ctxutil.CtxUtil,
) *AuthHandler {
	authHandler := &AuthHandler{
		router:         router,
		authService:    authService,
		authMiddleware: authMiddleware,
		ctxutil:        ctxUtil,
	}

	auth := router.Group("/auth")
	{
		auth.POST("/issue-token", authHandler.IssueToken)
		auth.POST("/validate", authHandler.Validate, authMiddleware.AuthenticateRequest)
	}

	return authHandler
}

// IssueToken
// @Summary IssueToken - 토큰 발급✅
// @Tags auth
// @Param IssueTokenRequest body presenter.IssueTokenRequest true
// @Produce json
// @Success 200 {object} presenter.IssueTokenResponse
// @Failure 404 {object} errors.Error "사용자를 찾을 수 없음 : user not found"
// @Router /auth/issue-token [post]
func (h AuthHandler) IssueToken(c echo.Context) error {
	var (
		request  presenter.IssueTokenRequest
		response *presenter.IssueTokenResponse
		err      error
	)

	if err = c.Bind(&request); err != nil {
		return err
	}

	response, err = h.authService.IssueToken(c.Request().Context(), request)

	switch {
	case err == nil:
		return c.JSON(http.StatusOK, response)
	case errors.Is(err, domain.ErrUserNotFound):
		return errors.Restore(err).SetStatus(http.StatusNotFound)
	default:
		return err
	}
}

// Validate
// @Summary Validate - 토큰 유효성 검사✅
// @Tags auth
// @Produce json
// @Success 200 {object} presenter.ValidateResponse
// @Failure 404 {object} errors.Error "사용자를 찾을 수 없음 : user not found"
// @Router /auth/validate [post]
// @Security BearerAuth
func (h AuthHandler) Validate(c echo.Context) error {
	var (
		request  presenter.ValidateRequest
		response *presenter.ValidateResponse
		err      error
	)

	request.RequestUserID, err = h.ctxutil.GetRequestUserID(c)
	if err != nil {
		return err
	}

	response, err = h.authService.Validate(c.Request().Context(), request)

	switch {
	case err == nil:
		return c.JSON(http.StatusOK, response)
	case errors.Is(err, domain.ErrUserNotFound):
		return errors.Restore(err).SetStatus(http.StatusNotFound)
	default:
		return err
	}
}
