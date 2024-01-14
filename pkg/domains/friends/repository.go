package friends

import (
	"context"
	"errors"
	"fmt"
	"social-media-app/pkg/constant"
	"social-media-app/pkg/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	AddFriend(ctx context.Context, friend *model.Friend) error
	GetFriends(ctx context.Context, userId uuid.UUID) ([]*model.Friend, error)
	GetFriendWithFriendId(ctx context.Context, friendId, userId uuid.UUID) (*model.Friend, error)
	AcceptFriend(ctx context.Context, friend *model.Friend) error
	DeleteFriend(ctx context.Context, friend *model.Friend) error
	RejectFriend(ctx context.Context, friend *model.Friend) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) AddFriend(ctx context.Context, friend *model.Friend) error {
	err := r.db.WithContext(ctx).Table("friends").Create(&friend).Error
	if err != nil {
		return errors.New(constant.FailedAddComment)
	}
	return nil
}

func (r *repository) GetFriends(ctx context.Context, userId uuid.UUID) ([]*model.Friend, error) {
	var friend []*model.Friend
	err := r.db.WithContext(ctx).Table("friends").Where("user_id = ? and status = 2", userId).Find(&friend).Error
	if err != nil {
		return friend, errors.New(constant.FailedAddComment)
	}
	return friend, nil
}

func (r *repository) GetFriendWithFriendId(ctx context.Context, friendId, userId uuid.UUID) (*model.Friend, error) {
	var friend *model.Friend
	err := r.db.WithContext(ctx).Where("user_id = ? and friend_id = ?", userId, friendId).First(&friend).Error
	if err != nil {
		return friend, errors.New(constant.FailedGetFriend)
	}
	return friend, nil
}

func (r *repository) AcceptFriend(ctx context.Context, friend *model.Friend) error {
	return r.db.Save(&friend).Error
}

func (r *repository) RejectFriend(ctx context.Context, friend *model.Friend) error {
	return r.db.Save(&friend).Error
}

func (r *repository) DeleteFriend(ctx context.Context, friend *model.Friend) error {
	err := r.db.Save(&friend).Error
	if err != nil {
		errMessage := fmt.Sprintf(constant.DeleteFailed, "friend")
		return errors.New(errMessage)
	}
	return r.db.Delete(&friend).Error
}
