package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	pkgErrors "example.com/golang-web/internal/pkg/errors"
	"example.com/golang-web/internal/pkg/response"
	"example.com/golang-web/internal/services"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=32"`
	Email    string `json:"email" binding:"required,email,max=128"`
	Password string `json:"password" binding:"required,min=6,max=128"`
}

type LoginRequest struct {
	Identifier string `json:"identifier" binding:"required,max=128"`
	Password   string `json:"password" binding:"required,min=1,max=128"`
}

type AuthHandler struct {
	authSvc services.AuthService
}

func NewAuthHandler(authSvc services.AuthService) *AuthHandler {
	return &AuthHandler{authSvc: authSvc}
}

func (h *AuthHandler) Register(api *gin.RouterGroup) {
	api.POST("/auth/register", h.register)
	api.POST("/auth/login", h.login)
}

func (h *AuthHandler) register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.AppErrorResponse(c, pkgErrors.NewAppError("VALIDATION_ERROR", "invalid register request", http.StatusBadRequest))
		return
	}

	accessToken, appErr := h.authSvc.Register(c.Request.Context(), services.RegisterInput{
		Username: strings.TrimSpace(req.Username),
		Email:    strings.TrimSpace(req.Email),
		Password: req.Password,
	})
	if appErr != nil {
		response.AppErrorResponse(c, appErr)
		return
	}

	response.Ok(c, http.StatusOK, "OK", "registered", gin.H{
		"accessToken": accessToken,
	})
}

func (h *AuthHandler) login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.AppErrorResponse(c, pkgErrors.NewAppError("VALIDATION_ERROR", "invalid login request", http.StatusBadRequest))
		return
	}

	accessToken, appErr := h.authSvc.Login(c.Request.Context(), services.LoginInput{
		Identifier: strings.TrimSpace(req.Identifier),
		Password:   req.Password,
	})
	if appErr != nil {
		response.AppErrorResponse(c, appErr)
		return
	}

	response.Ok(c, http.StatusOK, "OK", "logged in", gin.H{
		"accessToken": accessToken,
	})
}

