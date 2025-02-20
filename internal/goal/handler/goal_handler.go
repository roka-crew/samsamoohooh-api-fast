package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/roka-crew/internal/goal/service"
	"github.com/roka-crew/pkg/ctxutil"
	"github.com/roka-crew/presenter"
)

type GoalHandler struct {
	goalService *service.GoalService
	ctxutil     *ctxutil.CtxUtil
}

func NewGoalHandler(
	ctxutil *ctxutil.CtxUtil,
	goalService *service.GoalService,
) *GoalHandler {
	goalHandler := &GoalHandler{
		ctxutil:     ctxutil,
		goalService: goalService,
	}

	return goalHandler
}

// CreateGoal
// @Summary Create a goal - 목표 단건 생성 및 참여 ✅
// @Tags goals
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param CreateGoalRequest body presenter.CreateGoalRequest true "CreateGoalRequest"
// @Success 201 {object} presenter.CreateGoalResponse
// @Router /goals [post]
// @Security Bearer
func (h GoalHandler) CreateGoal(c echo.Context) error {
	var (
		request  presenter.CreateGoalRequest
		response presenter.CreateGoalResponse
		err      error
	)

	if err = c.Bind(&request); err != nil {
		return err
	}

	response, err = h.goalService.CreateGoal(c.Request().Context(), request)

	switch {
	case err != nil:
		return c.JSON(http.StatusCreated, response)
	default:
		return c.JSON(http.StatusCreated, response)
	}
}
