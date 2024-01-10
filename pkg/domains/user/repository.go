package user

import (
	"context"
	"social-media-app/pkg/model"

	"gorm.io/gorm"
)

type Repository interface {
	Save(ctx context.Context, user *model.User) error
	GetUserWithPhone(ctx context.Context, phone string) (model.User, error)
	GetUserWithEmail(ctx context.Context, email string) (model.User, error)
	GetUserWithUserName(ctx context.Context, username string) (model.User, error)
	FindUserWithIdentifiers(ctx context.Context, identifier string) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, user *model.User) error {
	return r.db.Create(&user).Error
}

func (r *repository) GetUserWithPhone(ctx context.Context, phone string) (model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Table("users").Where("phone = ?", phone).First(&user).Error
	return user, err
}

func (r *repository) GetUserWithEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Table("users").Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *repository) GetUserWithUserName(ctx context.Context, username string) (model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Table("users").Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *repository) FindUserWithIdentifiers(ctx context.Context, identifier string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ? OR username = ? OR phone = ?", identifier, identifier, identifier).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
