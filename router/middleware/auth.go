package middleware

import (
	"net/http"
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

func NewAuthMiddleware(
	token *token.Token,
	ctxutil *ctxutil.CtxUtil,
) *AuthMiddleware {
	return &AuthMiddleware{
		token:   token,
		ctxutil: ctxutil,
	}
}

func (m AuthMiddleware) AuthenticateRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return domain.ErrAuthorizationHeaderNeeded.
				SetStatus(http.StatusUnauthorized).
				SetDetail("Authorization 헤더가 필요합니다.")
		}

		// Bearer 토큰인지 확인
		if !strings.HasPrefix(tokenString, "Bearer ") {
			return domain.ErrInvalidTokenRequired.
				SetStatus(http.StatusUnauthorized).
				SetDetail("Bearer 토큰이 필요합니다.")
		}

		// Bearer 토큰에서 토큰 값만 추출
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		payload, err := m.token.ParseToken(tokenString)
		if err != nil {
			return domain.ErrInvalidToken.
				SetStatus(http.StatusUnauthorized).
				SetDetail("유효하지 않은 토큰입니다.")
		}

		m.ctxutil.SetTokenUser(c, payload)
		return next(c)
	}
}
