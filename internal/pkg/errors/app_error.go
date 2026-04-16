package errors

import "net/http"

// AppError 是业务层/接口层统一错误载体。
// Code 用于前端/调用方做可识别分支；Message 用于展示。
type AppError struct {
	Code       string
	Message    string
	HTTPStatus int
}

func NewAppError(code, message string, httpStatus int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: httpStatus,
	}
}

func (e *AppError) ResponseStatus() int {
	if e.HTTPStatus != 0 {
		return e.HTTPStatus
	}
	return http.StatusInternalServerError
}
