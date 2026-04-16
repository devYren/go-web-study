package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	pkgErrors "example.com/golang-web/internal/pkg/errors"
	"example.com/golang-web/internal/pkg/response"
)

// AuthStub 在第一阶段用于让路由/链路跑通。
// 后续 t5-t6 会替换为真正的 JWT 校验逻辑。
func AuthStub() gin.HandlerFunc {
	return func(c *gin.Context) {
		response.AppErrorResponse(c, pkgErrors.NewAppError("UNAUTHORIZED", "jwt required", http.StatusUnauthorized))
		c.Abort()
	}
}

