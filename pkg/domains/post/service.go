package post

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"social-media-app/pkg/constant"
	"social-media-app/pkg/dtos"
	"social-media-app/pkg/model"
	"social-media-app/pkg/redis"
	"social-media-app/pkg/utils"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	CreatePost(ctx context.Context, req *dtos.AddPostDTO, userId uuid.UUID) error
	UpdatePost(ctx context.Context, req *dtos.UpdatePostDTO, userId uuid.UUID) error
	DeletePost(ctx context.Context, userId, postId uuid.UUID) error
	GetUserAllPosts(ctx context.Context, userId uuid.UUID) ([]*model.Post, error)
	GetAllPosts(ctx context.Context) ([]*model.Post, error)
	GetPostsWithPostId(ctx context.Context, userId, postId uuid.UUID) ([]*model.Post, error)
	GetWithShortLink(ctx context.Context, shortLink string) (*model.Post, error)
}

type service struct {
	repository Repository
}

func NewPost(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) CreatePost(ctx context.Context, req *dtos.AddPostDTO, userId uuid.UUID) error {
	shortLink, err := utils.GenerateShortLink()

	if err != nil {
		return err
	}

	post := &model.Post{
		UserID:    userId,
		ShortLink: shortLink,
		Caption:   req.Caption,
	}

	err = s.repository.CreatePost(ctx, post)
	if err != nil {
		return err
	}
	// TODO: 5 dk içinde güncelelme hakkı
	key := fmt.Sprintf(constant.RedisWithPost, post.UserID, post.ID)
	bytePost, err := json.Marshal(post)
	if err != nil {
		return errors.New("failed to marshal post for redis")
	}
	err = redis.Set(ctx, key, string(bytePost), 5*time.Minute)
	if err != nil {
		return errors.New("failed to set redis")
	}
	return nil
}

func (s *service) UpdatePost(ctx context.Context, req *dtos.UpdatePostDTO, userId uuid.UUID) error {
	fmt.Println("service girdi")
	key := fmt.Sprintf(constant.RedisWithPost, userId, req.Id)
	isExist, _ := redis.Exists(ctx, key)
	if !isExist {
		return errors.New("timeout for updating post")
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return errors.New("uuid parse error")
	}
	post, err := s.repository.GetPostWithId(ctx, userId, id)
	if err != nil {
		errMessage := fmt.Sprintf(constant.UpdateFailed, "Post")
		return errors.New(errMessage)
	}

	post.MapperForUpdate(req)
	err = s.repository.UpadatePost(ctx, post)
	if err != nil {
		return err
	}
	fmt.Println("servis çıktı")
	return nil
}

func (s *service) DeletePost(ctx context.Context, userId, postId uuid.UUID) error {
	var (
		post *model.Post
		err  error
	)
	post, err = s.repository.GetPostWithId(ctx, userId, postId)
	if err != nil {
		return err
	}
	err = s.repository.DeletePost(ctx, post)
	if err != nil {
		return err
	}
	return err
}

func (s *service) GetUserAllPosts(ctx context.Context, userId uuid.UUID) ([]*model.Post, error) {
	posts, err := s.repository.GetUserAllPosts(ctx, userId)
	if err != nil {
		return posts, err
	}
	return posts, nil
}

func (s *service) GetAllPosts(ctx context.Context) ([]*model.Post, error) {
	posts, err := s.repository.GetAllPosts(ctx)
	if err != nil {
		return posts, err
	}
	return posts, nil
}

func (s *service) GetPostsWithPostId(ctx context.Context, userId, postId uuid.UUID) ([]*model.Post, error) {
	posts, err := s.repository.GetPostWithPostId(ctx, userId, postId)
	if err != nil {
		return posts, err
	}
	return posts, nil
}

func (s *service) GetWithShortLink(ctx context.Context, shortLink string) (*model.Post, error) {
	post, err := s.repository.GetPostWithShortLink(ctx, shortLink)
	if err != nil {
		return post, err
	}
	return post, nil
}
