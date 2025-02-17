package router

import (
	"github.com/labstack/echo/v4"
	"github.com/roka-crew/router/validator"
)

type Router struct {
	*echo.Echo
}

func New() *Router {
	r := Router{
		Echo: echo.New(),
	}

	r.Validator = validator.New()

	return &r
}
