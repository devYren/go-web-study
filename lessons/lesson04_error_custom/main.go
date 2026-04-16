package main

import (
	"errors"
	"fmt"
	"net/http"
)

// ==========================================
// 1. 哨兵错误（Sentinel Error）
// ==========================================
// 预定义的全局错误变量，用 errors.Is 判断
// 对标 Java 中自定义异常类的用法

var ErrNotFound = errors.New("not found")
var ErrUnauthorized = errors.New("unauthorized")

func findUser(id int) (string, error) {
	if id <= 0 {
		return "", ErrNotFound
	}
	return "Alice", nil
}

// ==========================================
// 2. 自定义错误类型 — 对标 Java 的自定义 Exception
// ==========================================

// Java:
// public class AppException extends RuntimeException {
//     private String code;
//     private int httpStatus;
//     public AppException(String code, String message, int status) { ... }
// }

// Go: 任何实现了 Error() string 方法的类型就是 error
type AppError struct {
	Code       string
	Message    string
	HTTPStatus int
}

// 实现 error 接口（error 接口只有一个方法：Error() string）
func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func NewAppError(code, message string, status int) *AppError {
	return &AppError{Code: code, Message: message, HTTPStatus: status}
}

// 模拟一个会返回 AppError 的业务函数
func login(username, password string) error {
	if username == "" {
		return NewAppError("VALIDATION_ERROR", "username is required", http.StatusBadRequest)
	}
	if password != "123456" {
		return NewAppError("INVALID_CREDENTIALS", "wrong password", http.StatusUnauthorized)
	}
	return nil
}

func main() {
	// ==========================================
	// errors.Is — 判断错误是不是某个特定错误
	// 对标 Java: catch (NotFoundException e)
	// ==========================================
	_, err := findUser(0)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("用户不存在") // ✅ 命中
	}

	// 即使错误被包装过，errors.Is 也能穿透
	wrappedErr := fmt.Errorf("findUser failed: %w", ErrNotFound)
	if errors.Is(wrappedErr, ErrNotFound) {
		fmt.Println("包装后依然能识别") // ✅ 命中
	}

	// ==========================================
	// errors.As — 提取特定类型的错误
	// 对标 Java: catch (AppException e) { e.getCode(); }
	// ==========================================
	err = login("admin", "wrong")
	// 想从 error 中取出 AppError 来读 Code 和 HTTPStatus
	var appErr *AppError
	if errors.As(err, &appErr) {
		fmt.Printf("业务错误: code=%s, status=%d, msg=%s\n",
			appErr.Code, appErr.HTTPStatus, appErr.Message)
	}

	// ==========================================
	// 对照你项目中的实际用法
	// ==========================================
	// 你项目里的 auth_service_impl.go 就用了类似的模式：
	//
	// var mysqlErr *mysqldriver.MySQLError
	// if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
	//     // MySQL 唯一索引冲突
	// }
	//
	// 这里 errors.As 把 error 提取为具体的 MySQLError 类型
	// 然后就能读取 MySQL 特有的错误码（1062 = duplicate entry）

	// ==========================================
	// 总结对比
	// ==========================================
	// Java:
	//   try { ... }
	//   catch (NotFoundException e) { ... }      → Go: errors.Is(err, ErrNotFound)
	//   catch (AppException e) { e.getCode(); }  → Go: errors.As(err, &appErr)
	//   throw new AppException("msg", cause);    → Go: fmt.Errorf("msg: %w", err)
}
