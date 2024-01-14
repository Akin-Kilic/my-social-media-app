package user

import (
	"context"
	"errors"
	"social-media-app/pkg/constant"
	"social-media-app/pkg/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Save(ctx context.Context, user *model.User) error
	GetUserWithId(ctx context.Context, userId uuid.UUID) (model.User, error)
	GetUserWithPhone(ctx context.Context, phone string) (model.User, error)
	GetUserWithEmail(ctx context.Context, email string) (model.User, error)
	GetUserWithUserName(ctx context.Context, username string) (model.User, error)
	FindUserWithIdentifiers(ctx context.Context, identifier string) (model.User, error)
	UpdateProfilePhoto(ctx context.Context, image *model.Image) error
	SaveImage(image *model.Image) error
	GetProfilePhoto(ctx context.Context, userId uuid.UUID) (*model.Image, error)
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

func (r *repository) GetUserWithId(ctx context.Context, userId uuid.UUID) (model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Table("users").Where("id = ?", userId).First(&user).Error
	return user, err
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

func (r *repository) UpdateProfilePhoto(ctx context.Context, image *model.Image) error {
	userId := image.UserId
	err := r.db.WithContext(ctx).Table("images").Where("user_id = ? and image_type = ?", userId, "2").Save(&image).Error
	if err != nil {
		return errors.New(constant.FailedUpdateUser)
	}
	return nil
}

func (r *repository) SaveImage(image *model.Image) error {
	return r.db.Create(image).Error
}

func (r *repository) GetProfilePhoto(ctx context.Context, userId uuid.UUID) (*model.Image, error) {
	var image *model.Image
	err := r.db.WithContext(ctx).Table("images").Where("user_id = ? and image_type = ?", userId, "1").First(&image).Error
	if err != nil {
		return image, errors.New(constant.FailedGetProfilePhoto)
	}
	return image, nil
}
