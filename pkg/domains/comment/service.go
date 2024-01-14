package comment

import (
	"context"
	"errors"
	"fmt"
	"social-media-app/pkg/constant"
	"social-media-app/pkg/dtos"
	"social-media-app/pkg/model"

	"github.com/google/uuid"
)

type Service interface {
	AddComment(ctx context.Context, req *model.Comment) error
	GetCommentsForPost(ctx context.Context, postId uuid.UUID) ([]model.Comment, error)
	UpdateComment(ctx context.Context, req *dtos.UpdateCommentDTO) error
	DeletePost(ctx context.Context, commentId uuid.UUID) error
}

type service struct {
	repository Repository
}

func NewComment(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) AddComment(ctx context.Context, req *model.Comment) error {
	return s.repository.AddComment(ctx, req)
}

func (s *service) GetCommentsForPost(ctx context.Context, postId uuid.UUID) ([]model.Comment, error) {
	comments, err := s.repository.GetCommentsForPost(ctx, postId)
	if err != nil {
		return comments, errors.New(constant.FailedGetPost)
	}
	return comments, nil
}

func (s *service) UpdateComment(ctx context.Context, req *dtos.UpdateCommentDTO) error {
	uuidComId, err := uuid.Parse(string(req.Id))
	if err != nil {
		return errors.New("parse uuid error")
	}
	comment, err := s.repository.GetCommentWithCommentId(ctx, uuidComId)
	if err != nil {
		errMessage := fmt.Sprintf(constant.UpdateFailed, "Post")
		return errors.New(errMessage)
	}

	comment.Mapper(req)
	err = s.repository.UpdateComment(ctx, comment)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeletePost(ctx context.Context, commentId uuid.UUID) error {
	var (
		comment *model.Comment
		err     error
	)
	comment, err = s.repository.GetCommentWithCommentId(ctx, commentId)
	if err != nil {
		return err
	}
	fmt.Println(comment)
	err = s.repository.DeleteComment(ctx, comment)
	if err != nil {
		return err
	}
	return err
}
