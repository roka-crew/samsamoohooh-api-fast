package handler

import (
	"net/http"

	"github.com/roka-crew/pkg/errors"

	"github.com/labstack/echo/v4"
	"github.com/roka-crew/domain"
	"github.com/roka-crew/internal/group/service"
	"github.com/roka-crew/pkg/ctxutil"
	"github.com/roka-crew/presenter"
	"github.com/roka-crew/router"
	"github.com/roka-crew/router/middleware"
)

type GroupHandler struct {
	ctxutil      *ctxutil.CtxUtil
	router       *router.Router
	groupService *service.GroupService
}

func NewGroupHandler(
	ctxutil *ctxutil.CtxUtil,
	router *router.Router,
	groupService *service.GroupService,
	authMiddleware *middleware.AuthMiddleware,
) *GroupHandler {
	groupHandler := &GroupHandler{
		ctxutil:      ctxutil,
		router:       router,
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

	request.RequestUserID, err = h.ctxutil.GetRequestUserID(c)
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
// @Summary List groups - 그룹 다건 조회 ✅
// @Tags groups
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} presenter.ListGroupsResponse
// @Failure 404 {object} errors.Error "그룹을 찾지 못함: group not found"
// @Router /groups [get]
// @Security Bearer
func (h GroupHandler) ListGroups(c echo.Context) error {
	var (
		request  presenter.ListGroupsRequest
		response presenter.ListGroupsResponse
		err      error
	)

	request.RequestUserID, err = h.ctxutil.GetRequestUserID(c)
	if err != nil {
		return err
	}

	response, err = h.groupService.ListGroups(c.Request().Context(), request)

	switch {
	case err == nil:
		return c.JSON(http.StatusOK, response)
	case errors.Is(err, domain.ErrGroupNotFound):
		return errors.Restore(err).SetStatus(http.StatusNotFound)
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
// @Failure 403 {object} errors.Error "그룹 소유자가 아님: group not owned"
// @Failure 404 {object} errors.Error "그룹을 찾지 못함: group not found"
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

	request.RequestUserID, err = h.ctxutil.GetRequestUserID(c)
	if err != nil {
		return err
	}

	err = h.groupService.PatchGroup(c.Request().Context(), request)

	switch {
	case err == nil:
		return c.NoContent(http.StatusNoContent)
	case errors.Is(err, domain.ErrGroupNotOwned):
		return errors.Restore(err).SetStatus(http.StatusForbidden).SetDetail("그룹 소유자가 아님")
	case errors.Is(err, domain.ErrGroupNotFound):
		return errors.Restore(err).SetStatus(http.StatusNotFound).SetDetail("그룹을 찾지 못함")
	default:
		return err
	}
}

// DeleteGroup
// @Summary Delete a group - 그룹 단건 삭제 ✅
// @Tags groups
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param group-id path string true "Group ID"
// @Success 204
// @Failure 403 {object} errors.Error "그룹 소유자가 아님: group not owned"
// @Failure 403 {object} errors.Error "그룹에 사용자가 있음: group has users"
// @Failure 404 {object} errors.Error "그룹을 찾지 못함: group not found"
// @Router /groups/{group-id} [delete]
// @Security Bearer
// func (h GroupHandler) DeleteGroup(c echo.Context) error {
// 	var (
// 		request presenter.DeleteGroupRequest
// 		err     error
// 	)

// 	if err := c.Bind(&request); err != nil {
// 		return err
// 	}

// 	request.RequestUserID, err = h.ctxutil.GetRequestUserID(c)
// 	if err != nil {
// 		return err
// 	}

// 	err = h.groupService.DeleteGroup(c.Request().Context(), request)

// 	switch {
// 	case err == nil:
// 		return c.NoContent(http.StatusNoContent)
// 	case errors.Is(err, domain.ErrGroupNotOwned):
// 		return errors.Restore(err).SetStatus(http.StatusForbidden).SetDetail("그룹 소유자가 아님")
// 	case errors.Is(err, domain.ErrGroupHasUsers):
// 		return errors.Restore(err).SetStatus(http.StatusForbidden).SetDetail("그룹에 사용자가 있음")
// 	case errors.Is(err, domain.ErrGroupNotFound):
// 		return errors.Restore(err).SetStatus(http.StatusNotFound).SetDetail("그룹을 찾지 못함")
// 	default:
// 		return err
// 	}
// }

// LeaveGroup
// @Summary Leave a group - 그룹 단건 탈퇴 ✅
// @Tags groups
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param group-id path string true "Group ID"
// @Success 204
// @Failure 403 {object} errors.Error "그룹 소유자가 아님: group not owned"
// @Failure 404 {object} errors.Error "그룹을 찾지 못함: group not found"
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

	request.RequestUserID, err = h.ctxutil.GetRequestUserID(c)
	if err != nil {
		return err
	}

	err = h.groupService.LeaveGroup(c.Request().Context(), request)

	switch {
	case err == nil:
		return c.NoContent(http.StatusNoContent)
	case errors.Is(err, domain.ErrGroupNotOwned):
		return errors.Restore(err).SetStatus(http.StatusForbidden).SetDetail("그룹 소유자가 아님")
	case errors.Is(err, domain.ErrGroupNotFound):
		return errors.Restore(err).SetStatus(http.StatusNotFound).SetDetail("그룹을 찾지 못함")
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
// @Failure 404 {object} errors.Error "그룹을 찾지 못함: group not found"
// @Failure 403 {object} errors.Error "그룹에 이미 참여함: group already joined"
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

	request.RequestUserID, err = h.ctxutil.GetRequestUserID(c)
	if err != nil {
		return err
	}

	err = h.groupService.JoinGroup(c.Request().Context(), request)

	switch {
	case err == nil:
		return c.NoContent(http.StatusNoContent)
	case errors.Is(err, domain.ErrGroupNotFound):
		return errors.Restore(err).SetStatus(http.StatusNotFound).SetDetail("그룹을 찾지 못함")
	case errors.Is(err, domain.ErrGroupAlreadyJoined):
		return errors.Restore(err).SetStatus(http.StatusForbidden).SetDetail("그룹에 이미 참여함")
	default:
		return err
	}
}
