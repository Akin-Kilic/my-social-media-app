package friends

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
	AddFriend(ctx context.Context, req *dtos.AddFriendDTO, userId uuid.UUID) error
	GetFriends(ctx context.Context, userId uuid.UUID) ([]*model.Friend, error)
	AcceptFriend(ctx context.Context, userId, friendId uuid.UUID) error
	RejectFriend(ctx context.Context, userId, friendId uuid.UUID) error
	DeleteFriend(ctx context.Context, userId, friendId uuid.UUID) error
}

type service struct {
	repository Repository
}

func NewFriends(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) AddFriend(ctx context.Context, req *dtos.AddFriendDTO, userId uuid.UUID) error {

	var friend model.Friend

	friend.UserId = userId
	friend.FriendId = req.FriendId
	fmt.Println(friend)

	err := s.repository.AddFriend(ctx, &friend)
	if err != nil {
		return errors.New("cannot save friends" + err.Error())
	}

	return nil
}

func (s *service) GetFriends(ctx context.Context, userId uuid.UUID) ([]*model.Friend, error) {
	friends, err := s.repository.GetFriends(ctx, userId)
	if err != nil {
		return friends, errors.New("cannot get friends")
	}
	return friends, nil
}

func (s *service) AcceptFriend(ctx context.Context, userId, friendId uuid.UUID) error {
	friend, err := s.repository.GetFriendWithFriendId(ctx, friendId, userId)
	if err != nil {
		return err
	}
	friend.Status = "2"
	err = s.repository.AcceptFriend(ctx, friend)
	if err != nil {
		errMessage := fmt.Sprintf(constant.UpdateFailed, "Post")
		return errors.New(errMessage)
	}
	return nil
}

func (s *service) RejectFriend(ctx context.Context, userId, friendId uuid.UUID) error {
	friend, err := s.repository.GetFriendWithFriendId(ctx, friendId, userId)
	if err != nil {
		return err
	}
	friend.Status = "3"
	err = s.repository.RejectFriend(ctx, friend)
	if err != nil {
		errMessage := fmt.Sprintf(constant.UpdateFailed, "friend")
		return errors.New(errMessage)
	}
	return nil
}

func (s *service) DeleteFriend(ctx context.Context, userId, friendId uuid.UUID) error {
	friend, err := s.repository.GetFriendWithFriendId(ctx, friendId, userId)
	if err != nil {
		return err
	}
	friend.Status = "4"
	err = s.repository.DeleteFriend(ctx, friend)
	if err != nil {
		errMessage := fmt.Sprintf(constant.UpdateFailed, "friend")
		return errors.New(errMessage)
	}
	return nil
}
