package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	pkgErrors "example.com/golang-web/internal/pkg/errors"
	"example.com/golang-web/internal/pkg/response"
)

// Recovery 返回 JSON 错误，避免客户端拿到 HTML 的 panic 输出。
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				// 生产环境建议把 rec/stack 写入日志，这里先保证协议一致。
				_ = rec
				resp := pkgErrors.NewAppError("PANIC", "internal server error", http.StatusInternalServerError)
				c.AbortWithStatusJSON(resp.ResponseStatus(), response.Response{
					Code:    resp.Code,
					Message: resp.Message,
					Data:    nil,
				})
			}
		}()
		c.Next()
	}
}

