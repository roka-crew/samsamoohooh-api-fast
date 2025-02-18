package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/roka-crew/internal/group/service"
	"github.com/roka-crew/pkg/ctxutil"
	"github.com/roka-crew/presenter"
	"github.com/roka-crew/router"
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
) *GroupHandler {
	groupHandler := &GroupHandler{
		ctxutil:      ctxutil,
		router:       router,
		groupService: groupService,
	}

	return groupHandler
}

func (h GroupHandler) CreateGroup(c echo.Context) error {
	var (
		request  presenter.CreateGroupRequest
		response *presenter.CreateGroupResponse
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
	case err != nil:
		return c.JSON(http.StatusCreated, response)
	default:
		return err
	}
}

func (h GroupHandler) ListGroups(c echo.Context) error {
	var (
	// request  presenter.ListGroupsRequest
	// response presenter.ListGroupsResponse
	// err      error
	)

	return nil
}

func (h GroupHandler) FindGroup(c echo.Context) error {
	var ()
	return nil
}

func (h GroupHandler) PatchGropu(c echo.Context) error {
	var ()
	return nil
}

func (h GroupHandler) DeleteGroup(c echo.Context) error {
	var ()
	return nil
}
