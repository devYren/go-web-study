package services

import (
	"context"

	pkgErrors "example.com/golang-web/internal/pkg/errors"
)

type RegisterInput struct {
	Username string
	Email    string
	Password string
}

type LoginInput struct {
	Identifier string // username or email
	Password   string
}

type AuthService interface {
	Register(ctx context.Context, in RegisterInput) (accessToken string, appErr *pkgErrors.AppError)
	Login(ctx context.Context, in LoginInput) (accessToken string, appErr *pkgErrors.AppError)
}

