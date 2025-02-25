package utils

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/roka-crew/pkg/apperr"
	"github.com/roka-crew/pkg/config"
)

type ErrorHandler struct {
	cfg config.Config
}

func NewErrorHandler(cfg *config.Config) *ErrorHandler {
	return &ErrorHandler{
		cfg: *cfg,
	}
}

func (h ErrorHandler) ErrorHandler(err error, c echo.Context) {

	switch err.(type) {
	case *apperr.Error:
		appError := apperr.Restore(err)
		appError.Instance = c.Path()

		_ = c.JSON(appError.Status, appError)
	case *apperr.InternalError:
		if h.cfg.Env == "dev" {
			fmt.Println(err.(*apperr.InternalError).Pretty())
		}

		_ = c.NoContent(http.StatusInternalServerError)
		return
	default:
		defaultHTTPErrorHandler(err, c)
	}
}

func defaultHTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	// Issue #1426
	code := he.Code
	message := he.Message

	// Send response
	if c.Request().Method == http.MethodHead { // Issue #608
		err = c.NoContent(he.Code)
	} else {
		err = c.JSON(code, message)
	}
}
