package comment

import (
	"context"
	"errors"
	"social-media-app/pkg/constant"
	"social-media-app/pkg/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	AddComment(ctx context.Context, comment *model.Comment) error
	GetCommentWithCommentId(ctx context.Context, commentId uuid.UUID) (*model.Comment, error)
	UpdateComment(ctx context.Context, comment *model.Comment) error
	DeleteComment(ctx context.Context, comment *model.Comment) error
	GetCommentsForPost(ctx context.Context, postId uuid.UUID) ([]model.Comment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) AddComment(ctx context.Context, comment *model.Comment) error {
	err := r.db.Table("comments").Create(&comment).Error
	if err != nil {
		return errors.New(constant.FailedAddComment)
	}
	return nil
}

func (r *repository) GetCommentWithCommentId(ctx context.Context, commentId uuid.UUID) (*model.Comment, error) {
	var comment *model.Comment
	err := r.db.WithContext(ctx).Table("comments").Where("id = ?", commentId).First(&comment).Error
	if err != nil {
		return comment, errors.New(constant.CommentNotFound)
	}
	return comment, nil
}

func (r *repository) UpdateComment(ctx context.Context, comment *model.Comment) error {
	err := r.db.WithContext(ctx).Table("comments").Save(&comment).Error
	if err != nil {
		return errors.New(constant.FailedUpdateComment)
	}
	return nil
}

func (r *repository) DeleteComment(ctx context.Context, comment *model.Comment) error {
	err := r.db.WithContext(ctx).Table("comments").Delete(&comment).Error
	if err != nil {
		return errors.New(constant.FailedDeleteComment)
	}
	return nil
}

func (r *repository) GetCommentsForPost(ctx context.Context, postId uuid.UUID) ([]model.Comment, error) {
	var comments []model.Comment
	err := r.db.WithContext(ctx).Where("post_id = ?", postId).Find(&comments).Error
	return comments, err
}
