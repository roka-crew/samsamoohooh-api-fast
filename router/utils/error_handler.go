package utils

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/roka-crew/pkg/config"
	"github.com/roka-crew/pkg/errors"
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
	case errors.Error:
		e := errors.Restore(err)
		e.Status = e.StatusCode()
		e.Title = http.StatusText(e.StatusCode())
		e.Instance = c.Path()

		if e.Status == http.StatusInternalServerError {
			if h.cfg.Env == "dev" {
				log.Printf("[ERROR] occured error: %+v\n", e)
			}
		}

		_ = c.JSON(e.StatusCode(), e)

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
