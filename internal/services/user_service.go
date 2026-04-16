package services

import (
	"context"

	"example.com/golang-web/internal/domain"
	pkgErrors "example.com/golang-web/internal/pkg/errors"
)

type UpdateUserInput struct {
	Username string
	Email    string
}

type UserService interface {
	GetByID(ctx context.Context, id uint, viewerID uint) (u *domain.User, appErr *pkgErrors.AppError)
	UpdateByID(ctx context.Context, id uint, viewerID uint, in UpdateUserInput) (u *domain.User, appErr *pkgErrors.AppError)
	DeleteByID(ctx context.Context, id uint, viewerID uint) *pkgErrors.AppError
}

