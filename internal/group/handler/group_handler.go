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

func (h GroupHandler) PatchGropu(c echo.Context) error {
	var ()
	return nil
}

func (h GroupHandler) DeleteGroup(c echo.Context) error {
	var ()
	return nil
}

func (h GroupHandler) LeaveGroup(c echo.Context) error {
	var ()
	return nil
}
