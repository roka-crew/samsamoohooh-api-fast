package router

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/roka-crew/pkg/errors"
)

var errorHandler = func(err error, c echo.Context) {
	switch err.(type) {
	case *echo.HTTPError:
		var message string

		echoError := err.(*echo.HTTPError)
		if m, ok := echoError.Message.(string); ok {
			message = m
		}

		_ = c.JSON(
			echoError.Code,
			errors.New("").
				SetStatus(echoError.Code).
				SetTitle(http.StatusText(echoError.Code)).
				SetDetail(message),
		)
		return

	case errors.Error:
		err := err.(errors.Error)
		if err.Status == http.StatusInternalServerError {
			_ = c.String(err.Status, "Internal Server Error")
			// TODO: log
			log.Printf("[ERROR] %+v\n", err)
			return
		}

		_ = c.JSON(err.Status, err)
		return
	}
}
