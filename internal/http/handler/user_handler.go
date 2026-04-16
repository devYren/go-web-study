package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	pkgErrors "example.com/golang-web/internal/pkg/errors"
	"example.com/golang-web/internal/pkg/response"
	"example.com/golang-web/internal/services"
)

type UpdateUserRequest struct {
	Username string `json:"username" binding:"omitempty,min=3,max=32"`
	Email    string `json:"email" binding:"omitempty,email,max=128"`
}

type UserHandler struct {
	userSvc services.UserService
}

func NewUserHandler(userSvc services.UserService) *UserHandler {
	return &UserHandler{userSvc: userSvc}
}

func (h *UserHandler) Register(apiUsers *gin.RouterGroup) {
	apiUsers.GET("/:id", h.getByID)
	apiUsers.PUT("/:id", h.updateByID)
	apiUsers.DELETE("/:id", h.deleteByID)
}

func viewerIDFromContext(c *gin.Context) (uint, bool) {
	v, ok := c.Get("viewerID")
	if !ok {
		return 0, false
	}
	id, ok := v.(uint)
	return id, ok
}

func (h *UserHandler) getByID(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.AppErrorResponse(c, pkgErrors.NewAppError("INVALID_ID", "invalid user id", http.StatusBadRequest))
		return
	}

	viewerID, ok := viewerIDFromContext(c)
	if !ok {
		response.AppErrorResponse(c, pkgErrors.NewAppError("UNAUTHORIZED", "jwt required", http.StatusUnauthorized))
		return
	}

	u, appErr := h.userSvc.GetByID(c.Request.Context(), uint(id64), viewerID)
	if appErr != nil {
		response.AppErrorResponse(c, appErr)
		return
	}

	response.Ok(c, http.StatusOK, "OK", "user found", u)
}

func (h *UserHandler) updateByID(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.AppErrorResponse(c, pkgErrors.NewAppError("INVALID_ID", "invalid user id", http.StatusBadRequest))
		return
	}

	viewerID, ok := viewerIDFromContext(c)
	if !ok {
		response.AppErrorResponse(c, pkgErrors.NewAppError("UNAUTHORIZED", "jwt required", http.StatusUnauthorized))
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.AppErrorResponse(c, pkgErrors.NewAppError("VALIDATION_ERROR", "invalid update request", http.StatusBadRequest))
		return
	}

	u, appErr := h.userSvc.UpdateByID(c.Request.Context(), uint(id64), viewerID, services.UpdateUserInput{
		Username: req.Username,
		Email:    req.Email,
	})
	if appErr != nil {
		response.AppErrorResponse(c, appErr)
		return
	}

	response.Ok(c, http.StatusOK, "OK", "user updated", u)
}

func (h *UserHandler) deleteByID(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.AppErrorResponse(c, pkgErrors.NewAppError("INVALID_ID", "invalid user id", http.StatusBadRequest))
		return
	}

	viewerID, ok := viewerIDFromContext(c)
	if !ok {
		response.AppErrorResponse(c, pkgErrors.NewAppError("UNAUTHORIZED", "jwt required", http.StatusUnauthorized))
		return
	}

	appErr := h.userSvc.DeleteByID(c.Request.Context(), uint(id64), viewerID)
	if appErr != nil {
		response.AppErrorResponse(c, appErr)
		return
	}

	response.Ok(c, http.StatusOK, "OK", "user deleted", gin.H{})
}

