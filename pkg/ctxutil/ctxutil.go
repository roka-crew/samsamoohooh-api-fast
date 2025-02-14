package ctxutil

import (
	"github.com/labstack/echo/v4"
	"github.com/roka-crew/pkg/token"
)

const (
	RequestUserPayload = "REQUEST_USER_PAYLOAD"
)

type CtxUtil struct {
}

func New() *CtxUtil {
	return &CtxUtil{}
}

func (CtxUtil) SetTokenUser(ctx echo.Context, payload *token.Payload) {
	ctx.Set(RequestUserPayload, payload)
}

func (CtxUtil) GetRequestUserID(ctx echo.Context) (uint, error) {
	payload, ok := ctx.Get(RequestUserPayload).(*token.Payload)
	if !ok {
		return 0, echo.NewHTTPError(500, "failed to get user from context")
	}

	return payload.UserID, nil
}
