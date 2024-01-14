package like

import (
	"context"
	"errors"
	"fmt"
	"social-media-app/pkg/constant"
	"social-media-app/pkg/model"

	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, req *model.Likes) error
	GetLikesWithPostId(ctx context.Context, userId uuid.UUID) ([]*model.Likes, error)
	GetLikesWithCommentId(ctx context.Context, commentId uuid.UUID) ([]*model.Likes, error)
	GetLikesCountForPost(ctx context.Context, postId string) (int64, error)
	DeleteLike(ctx context.Context, likeId string) error
}

type service struct {
	repository Repository
}

func NewLike(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) Create(ctx context.Context, req *model.Likes) error {
	fmt.Println("service girdi")
	return s.repository.CreateLike(*req)
}

func (s *service) GetLikesWithPostId(ctx context.Context, postId uuid.UUID) ([]*model.Likes, error) {
	posts, err := s.repository.GetLikesWithPostId(ctx, postId)
	if err != nil {
		return posts, err
	}
	return posts, nil
}

func (s *service) GetLikesWithCommentId(ctx context.Context, postId uuid.UUID) ([]*model.Likes, error) {
	posts, err := s.repository.GetLikesWithCommentId(ctx, postId)
	if err != nil {
		return posts, err
	}
	return posts, nil
}

func (s *service) GetLikesCountForPost(ctx context.Context, postId string) (int64, error) {
	count, err := s.repository.GetPostLikesCount(ctx, postId)
	if err != nil {
		return count, errors.New(constant.FailedGetPost)
	}
	return count, nil
}

func (s *service) DeleteLike(ctx context.Context, likeId string) error {
	var (
		like model.Likes
		err  error
	)
	like, err = s.repository.GetLikesWithId(ctx, likeId)
	if err != nil {
		return err
	}
	err = s.repository.DeleteLike(ctx, &like)
	if err != nil {
		return err
	}
	return err
}
