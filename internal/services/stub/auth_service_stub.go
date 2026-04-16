package stub

import (
	"context"

	pkgErrors "example.com/golang-web/internal/pkg/errors"
	"example.com/golang-web/internal/services"
)

type AuthServiceStub struct{}

func NewAuthService() services.AuthService {
	return &AuthServiceStub{}
}

func (s *AuthServiceStub) Register(ctx context.Context, in services.RegisterInput) (string, *pkgErrors.AppError) {
	return "", pkgErrors.NewAppError("NOT_IMPLEMENTED", "auth register not implemented yet", 501)
}

func (s *AuthServiceStub) Login(ctx context.Context, in services.LoginInput) (string, *pkgErrors.AppError) {
	return "", pkgErrors.NewAppError("NOT_IMPLEMENTED", "auth login not implemented yet", 501)
}
