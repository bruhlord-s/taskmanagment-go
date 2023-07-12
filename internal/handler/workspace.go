package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/bruhlord-s/openboard-go/internal/model"
	"github.com/gin-gonic/gin"
)

// @Summary Create workspace
// @Security ApiKeyAuth
// @Tags workspace
// @Description create workspace
// @ID create-workspace
// @Accept json
// @Produce json
// @Param input body model.Workspace true "workspace info"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/workspace [post]
func (h *Handler) createWorkspace(c *gin.Context) {
	userId, err := h.services.Authorization.GetUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input model.Workspace
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Workspace.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllWorkspacesResponse struct {
	Data []model.Workspace `json:"data"`
}

// @Summary Get all user's workspaces
// @Security ApiKeyAuth
// @Tags workspace
// @Description get workspaces
// @ID get-workspaces
// @Accept json
// @Produce json
// @Success 200 {object} []model.Workspace
// @Failure 500 {object} errorResponse
// @Router /api/v1/workspace [get]
func (h *Handler) getAllWorkspaces(c *gin.Context) {
	userId, err := h.services.Authorization.GetUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	workspaces, err := h.services.Workspace.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllWorkspacesResponse{
		Data: workspaces,
	})
}

// @Summary Get workspace by id
// @Security ApiKeyAuth
// @Tags workspace
// @Description get workspace
// @ID get-workspace
// @Accept json
// @Produce json
// @Success 200 {object} model.Workspace
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/workspace/{id} [get]
func (h *Handler) getWorkspaceById(c *gin.Context) {
	userId, err := h.services.Authorization.GetUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	workspace, err := h.services.Workspace.GetById(userId, id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			newErrorResponse(c, http.StatusNotFound, "no workspaces found with given id")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, workspace)
}

// @Summary Update workspace
// @Security ApiKeyAuth
// @Tags workspace
// @Description update workspace
// @ID update-workspace
// @Accept json
// @Produce json
// @Param input body model.UpdateWorkspaceInput true "update info"
// @Success 204
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/workspace/{id} [put]
func (h *Handler) updateWorkspace(c *gin.Context) {
	userId, err := h.services.Authorization.GetUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input model.UpdateWorkspaceInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Workspace.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Delete workspace
// @Security ApiKeyAuth
// @Tags workspace
// @Description delete workspace
// @ID delete-workspace
// @Accept json
// @Produce json
// @Success 204
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/workspace/{id} [delete]
func (h *Handler) deleteWorkspace(c *gin.Context) {
	userId, err := h.services.Authorization.GetUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Workspace.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
