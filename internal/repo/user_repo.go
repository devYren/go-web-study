package repo

import (
	"context"
	"errors"

	"example.com/golang-web/internal/model"
	"gorm.io/gorm"
)

// UserRepository 定义用户数据访问接口。
type UserRepository interface {
	Create(ctx context.Context, u *model.UserModel) error
	FindByID(ctx context.Context, id uint) (*model.UserModel, error)
	FindByUsername(ctx context.Context, username string) (*model.UserModel, error)
	FindByEmail(ctx context.Context, email string) (*model.UserModel, error)
	FindByUsernameOrEmail(ctx context.Context, identifier string) (*model.UserModel, error)
	Update(ctx context.Context, id uint, fields map[string]interface{}) error
	Delete(ctx context.Context, id uint) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, u *model.UserModel) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func (r *userRepo) FindByID(ctx context.Context, id uint) (*model.UserModel, error) {
	var u model.UserModel
	if err := r.db.WithContext(ctx).First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) FindByUsername(ctx context.Context, username string) (*model.UserModel, error) {
	var u model.UserModel
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*model.UserModel, error) {
	var u model.UserModel
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) FindByUsernameOrEmail(ctx context.Context, identifier string) (*model.UserModel, error) {
	var u model.UserModel
	if err := r.db.WithContext(ctx).Where("username = ? OR email = ?", identifier, identifier).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) Update(ctx context.Context, id uint, fields map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.UserModel{}).Where("id = ?", id).Updates(fields).Error
}

func (r *userRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.UserModel{}, id).Error
}

func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
