package handler

import (
	"net/http"

	"github.com/roka-crew/pkg/apperr"

	"github.com/labstack/echo/v4"
	"github.com/roka-crew/domain"
	"github.com/roka-crew/internal/group/service"
	"github.com/roka-crew/pkg/ctxutil"
	"github.com/roka-crew/presenter"
	"github.com/roka-crew/router"
	"github.com/roka-crew/router/middleware"
)

type GroupHandler struct {
	groupService *service.GroupService
	ctxUtil      *ctxutil.CtxUtil
}

func NewGroupHandler(
	groupService *service.GroupService,
	ctxUtil *ctxutil.CtxUtil,
	router *router.Router,
	authMiddleware *middleware.AuthMiddleware,
) *GroupHandler {
	groupHandler := &GroupHandler{
		ctxUtil:      ctxUtil,
		groupService: groupService,
	}

	groups := router.Group("/groups", authMiddleware.AuthenticateRequest)
	{
		groups.GET("", groupHandler.ListGroups)
		groups.POST("", groupHandler.CreateGroup)
		groups.PATCH("/:group-id", groupHandler.PatchGroup)
		// groups.DELETE("/:group-id", groupHandler.DeleteGroup)
		groups.POST("/:group-id/leave", groupHandler.LeaveGroup)
		groups.POST("/:group-id/join", groupHandler.JoinGroup)
	}

	return groupHandler
}

// CreateGroup
// @Summary Create a group - 그룹 단건 생성 및 참여 ✅
// @Tags groups
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param CreateGroupRequest body presenter.CreateGroupRequest true "CreateGroupRequest"
// @Success 201 {object} presenter.CreateGroupResponse
// @Router /groups [post]
// @Security Bearer
func (h GroupHandler) CreateGroup(c echo.Context) error {
	var (
		request  presenter.CreateGroupRequest
		response presenter.CreateGroupResponse
		err      error
	)

	if err = c.Bind(&request); err != nil {
		return err
	}

	request.RequestUserID, err = h.ctxUtil.GetRequestUserID(c)
	if err != nil {
		return err
	}

	response, err = h.groupService.CreateGroup(c.Request().Context(), request)

	switch {
	case err == nil:
		return c.JSON(http.StatusCreated, response)
	default:
		return err
	}
}

// ListGroups
// @Summary list groups - 그룹 다건 조회 ✅
// @Tags groups
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} presenter.ListGroupsResponse
// @Failure 404 {object} apperr.Error "그룹을 찾지 못함: group not found"
// @Router /groups [get]
// @Security Bearer
func (h GroupHandler) ListGroups(c echo.Context) error {
	var (
		request  presenter.ListGroupsRequest
		response presenter.ListGroupsResponse
		err      error
	)

	request.RequestUserID, err = h.ctxUtil.GetRequestUserID(c)
	if err != nil {
		return err
	}

	response, err = h.groupService.ListGroups(c.Request().Context(), request)

	switch {
	case err == nil:
		return c.JSON(http.StatusOK, response)
	case apperr.Is(err, domain.ErrGroupNotFound):
		return apperr.Restore(err).SetStatus(http.StatusNotFound)
	default:
		return err
	}
}

// PatchGroup
// @Summary Patch a group - 그룹 단건 수정 ✅
// @Tags groups
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param group-id path string true "Group ID"
// @Param PatchGroupRequest body presenter.PatchGroupRequest true "PatchGroupRequest"
// @Success 204
// @Failure 403 {object} apperr.Error "그룹 소유자가 아님: group not member"
// @Failure 404 {object} apperr.Error "그룹을 찾지 못함: group not found"
// @Router /groups/{group-id} [patch]
// @Security Bearer
func (h GroupHandler) PatchGroup(c echo.Context) error {
	var (
		request presenter.PatchGroupRequest
		err     error
	)

	if err = c.Bind(&request); err != nil {
		return err
	}

	request.RequestUserID, err = h.ctxUtil.GetRequestUserID(c)
	if err != nil {
		return err
	}

	err = h.groupService.PatchGroup(c.Request().Context(), request)

	switch {
	case err == nil:
		return c.NoContent(http.StatusNoContent)
	case apperr.Is(err, domain.ErrGroupNotMember):
		return apperr.Restore(err).SetStatus(http.StatusForbidden)
	case apperr.Is(err, domain.ErrGroupNotFound):
		return apperr.Restore(err).SetStatus(http.StatusNotFound)
	default:
		return err
	}
}

// LeaveGroup
// @Summary Leave a group - 그룹 단건 탈퇴 ✅
// @Tags groups
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param group-id path string true "Group ID"
// @Success 204
// @Failure 403 {object} apperr.Error "그룹 소유자가 아님: group not member"
// @Failure 404 {object} apperr.Error "그룹을 찾지 못함: group not found"
// @Router /groups/{group-id}/leave [post]
// @Security Bearer
func (h GroupHandler) LeaveGroup(c echo.Context) error {
	var (
		request presenter.LeaveGroupRequest
		err     error
	)

	if err := c.Bind(&request); err != nil {
		return err
	}

	request.RequestUserID, err = h.ctxUtil.GetRequestUserID(c)
	if err != nil {
		return err
	}

	err = h.groupService.LeaveGroup(c.Request().Context(), request)

	switch {
	case err == nil:
		return c.NoContent(http.StatusNoContent)
	case apperr.Is(err, domain.ErrGroupNotMember):
		return apperr.Restore(err).SetStatus(http.StatusForbidden)
	case apperr.Is(err, domain.ErrGroupNotFound):
		return apperr.Restore(err).SetStatus(http.StatusNotFound)
	default:
		return err
	}
}

// JoinGroup
// @Summary Join a group - 그룹 단건 참여 ✅
// @Tags groups
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param group-id path string true "Group ID"
// @Success 204
// @Failure 404 {object} apperr.Error "그룹을 찾지 못함: group not found"
// @Failure 403 {object} apperr.Error "그룹에 이미 참여함: group already joined"
// @Router /groups/{group-id}/join [post]
// @Security Bearer
func (h GroupHandler) JoinGroup(c echo.Context) error {
	var (
		request presenter.JoinGroupRequest
		err     error
	)

	if err := c.Bind(&request); err != nil {
		return err
	}

	request.RequestUserID, err = h.ctxUtil.GetRequestUserID(c)
	if err != nil {
		return err
	}

	err = h.groupService.JoinGroup(c.Request().Context(), request)

	switch {
	case err == nil:
		return c.NoContent(http.StatusNoContent)
	case apperr.Is(err, domain.ErrGroupNotFound):
		return apperr.Restore(err).SetStatus(http.StatusNotFound)
	case apperr.Is(err, domain.ErrGroupAlreadyJoined):
		return apperr.Restore(err).SetStatus(http.StatusForbidden)
	default:
		return err
	}
}
