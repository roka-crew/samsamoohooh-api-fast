package main

import (
	authHandler "github.com/roka-crew/internal/auth/handler"
	authService "github.com/roka-crew/internal/auth/service"
	userHandler "github.com/roka-crew/internal/user/handler"
	userService "github.com/roka-crew/internal/user/service"
	userStore "github.com/roka-crew/internal/user/store"
	"github.com/roka-crew/pkg/config"
	"github.com/roka-crew/pkg/ctxutil"
	"github.com/roka-crew/pkg/persistence/sqlite"
	"github.com/roka-crew/pkg/token"
	"github.com/roka-crew/router"
	"github.com/roka-crew/router/middleware"
	"github.com/roka-crew/router/utils"
	"go.uber.org/fx"
)

var (
	tag = "v0.0.0"
)

func main() {
	fx.New(
		fx.Supply("configs/env.yaml"),
		fx.Provide(
			config.New,
			sqlite.New,
			utils.NewErrorHandler,
			token.New,
			ctxutil.New,

			middleware.NewAuthMiddleware,

			userHandler.NewUserHandler,
			userService.NewUserService,
			userStore.NewUserStore,

			authHandler.NewAuthHandler,
			authService.NewAuthService,

			router.New,
		),
		fx.Invoke(
			func(db *sqlite.SQLite) {},

			userHandler.NewUserHandler,
			authHandler.NewAuthHandler,

			func(r *router.Router) {},
		),
	).Run()
}
