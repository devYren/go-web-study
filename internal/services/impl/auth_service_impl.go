package impl

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"example.com/golang-web/internal/model"
	pkgErrors "example.com/golang-web/internal/pkg/errors"
	"example.com/golang-web/internal/repo"
	"example.com/golang-web/internal/services"
	mysqldriver "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	userRepo repo.UserRepository
}

func NewAuthService(userRepo repo.UserRepository) services.AuthService {
	return &AuthServiceImpl{userRepo: userRepo}
}

func (s *AuthServiceImpl) Register(ctx context.Context, in services.RegisterInput) (string, *pkgErrors.AppError) {
	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.ErrorContext(ctx, "bcrypt hash failed", "err", err)
		return "", pkgErrors.NewAppError("INTERNAL", "failed to hash password", http.StatusInternalServerError)
	}

	u := &model.UserModel{
		Username:     in.Username,
		Email:        in.Email,
		PasswordHash: string(hash),
	}
	if err := s.userRepo.Create(ctx, u); err != nil {
		if mysqlErr := asMySQLDuplicateEntry(err); mysqlErr != nil {
			return "", duplicateEntryToAppError(mysqlErr)
		}
		slog.ErrorContext(ctx, "failed to create user", "err", err)
		return "", pkgErrors.NewAppError("INTERNAL", "failed to create user", http.StatusInternalServerError)
	}

	// TODO: replace with real JWT in a later phase
	token := fmt.Sprintf("token-%d", u.ID)
	return token, nil
}

func (s *AuthServiceImpl) Login(ctx context.Context, in services.LoginInput) (string, *pkgErrors.AppError) {
	u, err := s.userRepo.FindByUsernameOrEmail(ctx, in.Identifier)
	if err != nil {
		if repo.IsNotFound(err) {
			return "", pkgErrors.NewAppError("INVALID_CREDENTIALS", "invalid username/email or password", http.StatusUnauthorized)
		}
		slog.ErrorContext(ctx, "failed to find user", "err", err)
		return "", pkgErrors.NewAppError("INTERNAL", "database error", http.StatusInternalServerError)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(in.Password)); err != nil {
		return "", pkgErrors.NewAppError("INVALID_CREDENTIALS", "invalid username/email or password", http.StatusUnauthorized)
	}

	token := fmt.Sprintf("token-%d", u.ID)
	return token, nil
}

// asMySQLDuplicateEntry 判断 err 是否为 MySQL 1062 (duplicate entry)。
func asMySQLDuplicateEntry(err error) *mysqldriver.MySQLError {
	var mysqlErr *mysqldriver.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return mysqlErr
	}
	return nil
}

// duplicateEntryToAppError 根据冲突的索引名返回具体的业务错误。
func duplicateEntryToAppError(mysqlErr *mysqldriver.MySQLError) *pkgErrors.AppError {
	msg := mysqlErr.Message
	switch {
	case strings.Contains(msg, "username"):
		return pkgErrors.NewAppError("USERNAME_TAKEN", "username already exists", http.StatusConflict)
	case strings.Contains(msg, "email"):
		return pkgErrors.NewAppError("EMAIL_TAKEN", "email already exists", http.StatusConflict)
	default:
		return pkgErrors.NewAppError("DUPLICATE", "duplicate entry", http.StatusConflict)
	}
}
