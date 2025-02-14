package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/roka-crew/pkg/ctxutil"
	"github.com/roka-crew/pkg/token"
)

type AuthMiddleware struct {
	ctxutil *ctxutil.CtxUtil
	token   *token.Token
}

func NewAuthMiddleware(token *token.Token) *AuthMiddleware {
	return &AuthMiddleware{token: token}
}

func (m AuthMiddleware) AuthenticateRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is required")
		}

		// Bearer 토큰인지 확인
		if !strings.HasPrefix(tokenString, "Bearer ") {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token: Bearer token is required")
		}

		payload, err := m.token.ParseToken(tokenString)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		m.ctxutil.SetTokenUser(c, payload)
		return nil
	}
}
