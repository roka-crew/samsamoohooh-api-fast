package main

import (
	authHandler "github.com/roka-crew/internal/auth/handler"
	authService "github.com/roka-crew/internal/auth/service"
	goalHandler "github.com/roka-crew/internal/goal/handler"
	goalService "github.com/roka-crew/internal/goal/service"
	goalStore "github.com/roka-crew/internal/goal/store"
	groupHandler "github.com/roka-crew/internal/group/handler"
	groupService "github.com/roka-crew/internal/group/service"
	groupStore "github.com/roka-crew/internal/group/store"
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

			groupHandler.NewGroupHandler,
			groupService.NewGroupService,
			groupStore.NewGroupStore,

			goalHandler.NewGoalHandler,
			goalService.NewGoalService,
			goalStore.NewGoalStore,

			authHandler.NewAuthHandler,
			authService.NewAuthService,

			router.New,
		),
		fx.Invoke(
			func(db *sqlite.SQLite) {},

			userHandler.NewUserHandler,
			groupHandler.NewGroupHandler,
			goalHandler.NewGoalHandler,
			authHandler.NewAuthHandler,

			func(r *router.Router) {},
		),

		// configs
		fx.NopLogger,
	).Run()
}
