package handler

import (
	"github.com/roka-crew/domain"
	"github.com/roka-crew/pkg/errors"
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
// @Failure 400 {object} errors.Error "deadline이 유효하지 않음 : invalid deadline"
// @Failure 403 {object} errors.Error "그룹에 속해있지 않음: group not member"k
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
	case errors.Is(err, domain.ErrInvalidDeadline):
		return errors.Restore(err).SetStatus(http.StatusBadRequest)
	case errors.Is(err, domain.ErrGroupNotMember):
		return errors.Restore(err).SetStatus(http.StatusForbidden).SetDetail("그룹에 속해있지 않습니다.")
	default:
		return c.JSON(http.StatusCreated, response)
	}
}

// ListGoals
// @Summary List goals - 목표 리스트 조회 ✅
// @Tags goals
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param ListGoalsRequest body presenter.ListGoalsRequest true "ListGoalsRequest"
// @Success 200 {object} presenter.ListGoalsResponse
// @Failure 403 {object} errors.Error "그룹에 속해있지 않음: group not member"
// @Router /goals [get]
// @Security Bearer
func (h GoalHandler) ListGoals(c echo.Context) error {
	var (
		request  presenter.ListGoalsRequest
		response presenter.ListGoalsResponse
		err      error
	)

	if err = c.Bind(&request); err != nil {
		return err
	}

	response, err = h.goalService.ListGoals(c.Request().Context(), request)

	switch {
	case err == nil:
		return c.JSON(http.StatusOK, response)
	case errors.Is(err, domain.ErrGroupNotMember):
		return errors.Restore(err).SetStatus(http.StatusForbidden)
	default:
		return err
	}
}
