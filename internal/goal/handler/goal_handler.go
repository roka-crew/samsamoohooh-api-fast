package handler

import (
	"github.com/roka-crew/domain"
	"github.com/roka-crew/pkg/errors"
	"github.com/roka-crew/router"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/roka-crew/internal/goal/service"
	"github.com/roka-crew/pkg/ctxutil"
	"github.com/roka-crew/presenter"
)

type GoalHandler struct {
	goalService *service.GoalService
	ctxUtil     *ctxutil.CtxUtil
}

func NewGoalHandler(
	ctxUtil *ctxutil.CtxUtil,
	goalService *service.GoalService,
	router *router.Router,
) *GoalHandler {
	goalHandler := &GoalHandler{
		ctxUtil:     ctxUtil,
		goalService: goalService,
	}

	groups := router.Group("/goals")
	{
		groups.POST("", goalHandler.CreateGoal)
		groups.GET("", goalHandler.ListGoals)
		groups.PATCH("/:goal-id", goalHandler.PatchGoal)
		groups.DELETE("/:goal-id", goalHandler.DeleteGoal)
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

	request.RequestUserID, err = h.ctxUtil.GetRequestUserID(c)
	if err != nil {
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

	request.RequestUserID, err = h.ctxUtil.GetRequestUserID(c)
	if err != nil {
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

// PatchGoal
// @Summary Patch a goal - 목표 단건 수정 ✅
// @Tags goals
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param goal-id path int true "Goal ID"
// @Param PatchGoalRequest body presenter.PatchGoalRequest true "PatchGoalRequest"
// @Success 204
// @Failure 403 {object} errors.Error "그룹에 속해있지 않음: group not member"
// @Router /goals/{goal-id} [patch]
// @Security Bearer
func (h GoalHandler) PatchGoal(c echo.Context) error {
	var (
		request presenter.PatchGoalRequest
		err     error
	)

	if err = c.Bind(&request); err != nil {
		return err
	}

	request.RequestUserID, err = h.ctxUtil.GetRequestUserID(c)
	if err != nil {
		return err
	}

	err = h.goalService.PatchGoal(c.Request().Context(), request)

	switch {
	case err == nil:
		return c.NoContent(http.StatusNoContent)
	case errors.Is(err, domain.ErrGroupNotMember):
		return errors.Restore(err).SetStatus(http.StatusForbidden).SetDetail("그룹에 속해있지 않습니다.")
	case errors.Is(err, domain.ErrGoalNotFound):
		return errors.Restore(err).SetStatus(http.StatusNotFound).SetDetail("목표를 찾을 수 없습니다.")
	case errors.Is(err, domain.ErrInvalidDeadline):
		return errors.Restore(err).SetStatus(http.StatusBadRequest)
	default:
		return err
	}
}

// DeleteGoal
// @Summary Delete a goal - 목표 단건 삭제 ✅
// @Tags goals
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param goal-id path int true "Goal ID"
// @Success 204
// @Failure 403 {object} errors.Error "그룹에 속해있지 않음: group not member"
// @Failure 404 {object} errors.Error "목표를 찾을 수 없음: goal not found"
// @Router /goals/{goal-id} [delete]
// @Security Bearer
func (h GoalHandler) DeleteGoal(e echo.Context) error {
	var (
		request presenter.DeleteGoalRequest
		err     error
	)

	if err = e.Bind(&request); err != nil {
		return err
	}

	request.RequestUserID, err = h.ctxUtil.GetRequestUserID(e)
	if err != nil {
		return err
	}

	err = h.goalService.DeleteGoal(e.Request().Context(), request)

	switch {
	case err == nil:
		return e.NoContent(http.StatusNoContent)
	case errors.Is(err, domain.ErrGroupNotMember):
		return errors.Restore(err).SetStatus(http.StatusForbidden).SetDetail("그룹에 속해있지 않습니다.")
	case errors.Is(err, domain.ErrGoalNotFound):
		return errors.Restore(err).SetStatus(http.StatusNotFound).SetDetail("목표를 찾을 수 없습니다.")
	default:
		return err
	}
}
