package post

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
	CreatePost(ctx context.Context, post *model.Post) error
	UpadatePost(ctx context.Context, post *model.Post) error
	DeletePost(ctx context.Context, post *model.Post) error
	GetPostWithId(ctx context.Context, userId, postId uuid.UUID) (*model.Post, error)
	GetUserAllPosts(ctx context.Context, userId uuid.UUID) ([]*model.Post, error)
	GetAllPosts(ctx context.Context) ([]*model.Post, error)
	GetPostWithPostId(ctx context.Context, userId, postId uuid.UUID) ([]*model.Post, error)
	GetPostWithShortLink(ctx context.Context, shortLink string) (*model.Post, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) CreatePost(ctx context.Context, post *model.Post) error {
	err := r.db.Table("posts").Save(&post).Error
	if err != nil {
		return errors.New(constant.FailedSharePost)
	}
	return nil
}

func (r *repository) UpadatePost(ctx context.Context, post *model.Post) error {
	fmt.Println("repository girdi")
	err := r.db.Table("posts").Save(post).Error
	if err != nil {
		return errors.New(constant.FailedUpdatePost)
	}
	fmt.Println("repository çıktı")
	return nil
}

func (r *repository) DeletePost(ctx context.Context, post *model.Post) error {
	err := r.db.Table("posts").Delete(&post).Error
	if err != nil {
		return errors.New(constant.FailedDeletePost)
	}
	return nil
}

func (r *repository) GetPostWithId(ctx context.Context, userId, postId uuid.UUID) (*model.Post, error) {
	var post *model.Post
	err := r.db.WithContext(ctx).Table("posts").Where("id = ? and user_id = ?", postId, userId).First(&post).Error
	if err != nil {
		return post, errors.New(constant.FailedGetPost)
	}
	return post, nil
}

func (r *repository) GetUserAllPosts(ctx context.Context, userId uuid.UUID) ([]*model.Post, error) {
	var posts []*model.Post
	err := r.db.WithContext(ctx).Table("posts").Where("user_id = ?", userId).Find(&posts).Error
	if err != nil {
		return posts, errors.New(constant.FailedGetPosts)
	}
	return posts, nil
}

func (r *repository) GetAllPosts(ctx context.Context) ([]*model.Post, error) {
	var posts []*model.Post
	err := r.db.WithContext(ctx).Table("posts").Find(&posts).Error
	if err != nil {
		return posts, errors.New(constant.FailedGetPosts)
	}
	return posts, nil
}

func (r *repository) GetPostWithPostId(ctx context.Context, userId, postId uuid.UUID) ([]*model.Post, error) {
	var posts []*model.Post
	err := r.db.WithContext(ctx).Table("posts").Where("user_id = ? and post_id", userId, postId).Find(&posts).Error
	if err != nil {
		return posts, errors.New(constant.FailedGetPosts)
	}
	return posts, nil
}

func (r *repository) GetPostWithShortLink(ctx context.Context, shortLink string) (*model.Post, error) {
	var post *model.Post
	err := r.db.WithContext(ctx).Table("posts").Where("short_link = ?", shortLink).First(&post).Error
	if err != nil {
		return post, errors.New(constant.FailedGetPosts)
	}
	return post, nil
}
