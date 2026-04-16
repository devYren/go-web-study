package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/golang-web/internal/http/handler"
	"example.com/golang-web/internal/http/middleware"
	"example.com/golang-web/internal/pkg/response"
	"example.com/golang-web/internal/services"
)

type RouterDeps struct {
	AuthSvc services.AuthService
	UserSvc services.UserService
}

func NewRouter(authSvc services.AuthService, userSvc services.UserService) *gin.Engine {
	r := gin.New()
	r.Use(middleware.Recovery())

	r.NoRoute(func(c *gin.Context) {
		// 统一走 JSON 协议，避免返回 HTML。
		response.Ok(c, http.StatusNotFound, "NOT_FOUND", "not found", nil)
	})

	health := handler.NewHealthHandler()
	health.Register(r)

	api := r.Group("/api/v1")

	authHandler := handler.NewAuthHandler(authSvc)
	authHandler.Register(api)

	users := api.Group("/users")
	users.Use(middleware.AuthStub())
	userHandler := handler.NewUserHandler(userSvc)
	userHandler.Register(users)

	return r
}
