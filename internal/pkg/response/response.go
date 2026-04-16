package response

import (
	"net/http"

	"errors"

	pkgErrors "example.com/golang-web/internal/pkg/errors"
)

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func WriteJSON(c interface {
	JSON(int, interface{})
	Status(int)
}, status int, resp interface{}) {
	c.JSON(status, resp)
}

// Ok 统一成功响应（200/201 由调用方控制状态码）。
func Ok(c interface {
	JSON(int, interface{})
	Status(int)
}, httpStatus int, code, message string, data interface{}) {
	WriteJSON(c, httpStatus, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// AppErrorResponse 统一错误响应。
func AppErrorResponse(c interface {
	JSON(int, interface{})
	Status(int)
}, err *pkgErrors.AppError) {
	status := err.ResponseStatus()
	if status <= 0 {
		status = http.StatusInternalServerError
	}
	Ok(c, status, err.Code, err.Message, nil)
}

func UnknownError(c interface {
	JSON(int, interface{})
	Status(int)
}, err error) {
	_ = errors.Unwrap(err)
	AppErrorResponse(c, pkgErrors.NewAppError("UNKNOWN_ERROR", "internal error", http.StatusInternalServerError))
}

