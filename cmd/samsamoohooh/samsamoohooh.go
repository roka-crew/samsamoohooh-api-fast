package main

import (
	userHandler "github.com/roka-crew/internal/user/handler"
	userService "github.com/roka-crew/internal/user/service"
	userStore "github.com/roka-crew/internal/user/store"
	"github.com/roka-crew/pkg/config"
	"github.com/roka-crew/pkg/ctxutil"
	"github.com/roka-crew/pkg/persistence/sqlite"
	"github.com/roka-crew/pkg/token"
	"github.com/roka-crew/router"
	"github.com/roka-crew/router/middleware"
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
			token.New,
			ctxutil.New,

			middleware.NewAuthMiddleware,

			userHandler.NewUserHandler,
			userService.NewUserService,
			userStore.NewUserStore,

			router.New,
		),
		fx.Invoke(
			func(db *sqlite.SQLite) {},

			userHandler.NewUserHandler,

			func(r *router.Router) {},
		),
	).Run()
}
