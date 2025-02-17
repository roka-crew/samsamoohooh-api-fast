package router

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/roka-crew/pkg/config"
	"github.com/roka-crew/router/validator"
	"go.uber.org/fx"
)

type Router struct {
	cfg *config.Config
	*echo.Echo
}

func New(
	cfg *config.Config,
	lc fx.Lifecycle,
) *Router {
	r := Router{
		Echo: echo.New(),
	}

	// echo configuration
	r.HideBanner = true
	r.Validator = validator.New()

	// lifecycle
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := r.Start(cfg.Listen); err != nil {
					log.Panicf("failed to start server: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := r.Shutdown(ctx); err != nil {
				log.Panicf("failed to shutdown server: %v", err)
			}

			return nil
		},
	})

	return &r
}
