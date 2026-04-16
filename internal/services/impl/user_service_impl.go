package impl

import (
	"context"
	"log/slog"
	"net/http"

	"example.com/golang-web/internal/domain"
	pkgErrors "example.com/golang-web/internal/pkg/errors"
	"example.com/golang-web/internal/repo"
	"example.com/golang-web/internal/services"
)

type UserServiceImpl struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) services.UserService {
	return &UserServiceImpl{userRepo: userRepo}
}

func (s *UserServiceImpl) GetByID(ctx context.Context, id uint, viewerID uint) (*domain.User, *pkgErrors.AppError) {
	u, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		if repo.IsNotFound(err) {
			return nil, pkgErrors.NewAppError("NOT_FOUND", "user not found", http.StatusNotFound)
		}
		slog.ErrorContext(ctx, "failed to find user", "id", id, "err", err)
		return nil, pkgErrors.NewAppError("INTERNAL", "database error", http.StatusInternalServerError)
	}
	return &domain.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}, nil
}

func (s *UserServiceImpl) UpdateByID(ctx context.Context, id uint, viewerID uint, in services.UpdateUserInput) (*domain.User, *pkgErrors.AppError) {
	if id != viewerID {
		return nil, pkgErrors.NewAppError("FORBIDDEN", "cannot update another user", http.StatusForbidden)
	}

	if _, err := s.userRepo.FindByID(ctx, id); err != nil {
		if repo.IsNotFound(err) {
			return nil, pkgErrors.NewAppError("NOT_FOUND", "user not found", http.StatusNotFound)
		}
		slog.ErrorContext(ctx, "failed to find user for update", "id", id, "err", err)
		return nil, pkgErrors.NewAppError("INTERNAL", "database error", http.StatusInternalServerError)
	}

	fields := make(map[string]interface{})
	if in.Username != "" {
		fields["username"] = in.Username
	}
	if in.Email != "" {
		fields["email"] = in.Email
	}

	if len(fields) > 0 {
		if err := s.userRepo.Update(ctx, id, fields); err != nil {
			slog.ErrorContext(ctx, "failed to update user", "id", id, "err", err)
			return nil, pkgErrors.NewAppError("INTERNAL", "failed to update user", http.StatusInternalServerError)
		}
	}

	u, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to reload user after update", "id", id, "err", err)
		return nil, pkgErrors.NewAppError("INTERNAL", "database error", http.StatusInternalServerError)
	}

	return &domain.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}, nil
}

func (s *UserServiceImpl) DeleteByID(ctx context.Context, id uint, viewerID uint) *pkgErrors.AppError {
	if id != viewerID {
		return pkgErrors.NewAppError("FORBIDDEN", "cannot delete another user", http.StatusForbidden)
	}

	if _, err := s.userRepo.FindByID(ctx, id); err != nil {
		if repo.IsNotFound(err) {
			return pkgErrors.NewAppError("NOT_FOUND", "user not found", http.StatusNotFound)
		}
		slog.ErrorContext(ctx, "failed to find user for delete", "id", id, "err", err)
		return pkgErrors.NewAppError("INTERNAL", "database error", http.StatusInternalServerError)
	}

	if err := s.userRepo.Delete(ctx, id); err != nil {
		slog.ErrorContext(ctx, "failed to delete user", "id", id, "err", err)
		return pkgErrors.NewAppError("INTERNAL", "failed to delete user", http.StatusInternalServerError)
	}
	return nil
}
