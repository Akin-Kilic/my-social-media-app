package like

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
	CreateLike(likes model.Likes) error
	GetLikesWithPostId(ctx context.Context, postId uuid.UUID) ([]*model.Likes, error)
	GetLikesWithCommentId(ctx context.Context, commentId uuid.UUID) ([]*model.Likes, error)
	GetPostLikesCount(ctx context.Context, postId string) (int64, error)
	GetLikesWithId(ctx context.Context, likeId string) (model.Likes, error)
	DeleteLike(ctx context.Context, like *model.Likes) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateLike(likes model.Likes) error {
	fmt.Println("repo girdi")
	return r.db.Table("likes").Create(&likes).Error
}

func (r *repository) GetLikesWithPostId(ctx context.Context, postId uuid.UUID) ([]*model.Likes, error) {
	var likes []*model.Likes
	err := r.db.WithContext(ctx).Table("likes").Where("post_id = ?", postId).Find(&likes).Error
	if err != nil {
		return likes, errors.New(constant.FailedGetPosts)
	}
	return likes, nil
}

func (r *repository) GetLikesWithCommentId(ctx context.Context, commentId uuid.UUID) ([]*model.Likes, error) {
	var likes []*model.Likes
	err := r.db.WithContext(ctx).Table("likes").Where("comment_id = ?", commentId).Find(&likes).Error
	if err != nil {
		return likes, errors.New(constant.FailedGetPosts)
	}
	return likes, nil
}

func (r *repository) GetPostLikesCount(ctx context.Context, postId string) (int64, error) {
	var count int64
	err := r.db.Table("likes").Where("post_id = ?", postId).Count(&count).Error
	return count, err
}

func (r *repository) GetLikesWithId(ctx context.Context, likeId string) (model.Likes, error) {
	var like model.Likes
	err := r.db.WithContext(ctx).Table("likes").Where("id = ?", likeId).First(&like).Error
	if err != nil {
		return like, errors.New(constant.LikeNotFound)
	}
	fmt.Println("repo get id")
	return like, nil
}

func (r *repository) DeleteLike(ctx context.Context, like *model.Likes) error {
	err := r.db.WithContext(ctx).Table("likes").Delete(&like).Error
	if err != nil {
		return errors.New(constant.FailedDeleteLike)
	}
	fmt.Println("repo delete")
	return nil
}
