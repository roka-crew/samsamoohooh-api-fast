package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/roka-crew/domain"
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
			return domain.ErrAuthorizationHeaderNeeded
		}

		// Bearer 토큰인지 확인
		if !strings.HasPrefix(tokenString, "Bearer ") {
			return domain.ErrInvalidTokenRequired
		}

		payload, err := m.token.ParseToken(tokenString)
		if err != nil {
			return domain.ErrInvalidToken
		}

		m.ctxutil.SetTokenUser(c, payload)
		return nil
	}
}
