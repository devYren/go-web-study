package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/golang-web/internal/pkg/response"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Register(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		response.Ok(c, http.StatusOK, "OK", "ok", nil)
	})
}

