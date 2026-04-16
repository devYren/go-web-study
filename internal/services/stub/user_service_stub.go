package stub

import (
	"context"

	"example.com/golang-web/internal/domain"
	pkgErrors "example.com/golang-web/internal/pkg/errors"
	"example.com/golang-web/internal/services"
)

type UserServiceStub struct{}

func NewUserService() services.UserService {
	return &UserServiceStub{}
}

func (s *UserServiceStub) GetByID(ctx context.Context, id uint, viewerID uint) (*domain.User, *pkgErrors.AppError) {
	return nil, pkgErrors.NewAppError("NOT_IMPLEMENTED", "user get not implemented yet", 501)
}

func (s *UserServiceStub) UpdateByID(ctx context.Context, id uint, viewerID uint, in services.UpdateUserInput) (*domain.User, *pkgErrors.AppError) {
	return nil, pkgErrors.NewAppError("NOT_IMPLEMENTED", "user update not implemented yet", 501)
}

func (s *UserServiceStub) DeleteByID(ctx context.Context, id uint, viewerID uint) *pkgErrors.AppError {
	return pkgErrors.NewAppError("NOT_IMPLEMENTED", "user delete not implemented yet", 501)
}

