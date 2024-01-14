package user

import (
	"context"
	"encoding/base64"
	"errors"
	"log"
	"os"
	"social-media-app/pkg/config"
	"social-media-app/pkg/constant"
	"social-media-app/pkg/dtos"
	"social-media-app/pkg/model"
	"social-media-app/pkg/utils"

	"github.com/google/uuid"
)

type Service interface {
	Register(ctx context.Context, user *model.User) error
	Login(ctx context.Context, req *dtos.LoginReq) (dtos.LoginDTO, error)
	Logout(ctx context.Context, key string) error
	UpdateProfilePhoto(ctx context.Context, path string, imageType string, userId, entityId uuid.UUID) error
	SaveImage(path string, imageType string, userId, entityId uuid.UUID) error
	GetProfilePhoto(ctx context.Context, userId uuid.UUID) (*model.Image, error)
}

type service struct {
	repository Repository
}

func NewUser(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) Register(ctx context.Context, user *model.User) error {
	err := s.duplicateControl(ctx, user.Username, user.Phone, user.Email)
	if err != nil {
		return err
	}
	err = user.PassHash()
	if err != nil {
		return errors.New(constant.PassHashFailed)
	}
	log.Println("user password: ", user.Password)

	err = s.repository.Save(ctx, user)
	if err != nil {
		return errors.New(constant.FailedCreateUser)
	}
	return nil
}

func (s *service) Login(ctx context.Context, req *dtos.LoginReq) (dtos.LoginDTO, error) {
	var loginDto dtos.LoginDTO

	user, err := s.repository.FindUserWithIdentifiers(ctx, req.Identifier)
	if err != nil {
		return loginDto, errors.New(constant.FailedLogin)
	}

	isTrue := utils.PasswordControl(user.Password, req.Password)
	if !isTrue {
		return loginDto, errors.New(constant.FailedUserNameOrPass)
	}

	token, err := utils.GenerateJwt(user.ID, config.ReadValue().JwtSecret)
	if err != nil {
		return loginDto, errors.New(constant.FailedLogin)
	}

	// key := fmt.Sprintf(constant.RedisForJwt, token, user.ID)
	// redis.Set(ctx, key, token, time.Hour*time.Duration(config.ReadValue().JwtExpTime))

	loginDto = dtos.LoginDTO{
		ID:       user.ID,
		UserName: user.Username,
		Token:    token,
	}

	return loginDto, nil
}

func (s *service) Logout(ctx context.Context, key string) error {
	// err := redis.Delete(key)
	// if err != nil {
	// 	return errors.New(constant.FailedLogout)
	// }

	return nil
}

func (s *service) duplicateControl(ctx context.Context, username, phone, email string) error {
	_, err := s.repository.GetUserWithPhone(ctx, phone)
	if err == nil {
		return errors.New(constant.AlreadyExistsPhone)
	}
	_, err = s.repository.GetUserWithUserName(ctx, username)
	if err == nil {
		return errors.New(constant.AlreadyExistsUsername)
	}
	_, err = s.repository.GetUserWithEmail(ctx, email)
	if err == nil {
		return errors.New(constant.AlreadyExistsEmail)
	}
	return nil
}

func (s *service) UpdateProfilePhoto(ctx context.Context, path string, imageType string, userId, entityId uuid.UUID) error {
	imageBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	base64Image := base64.StdEncoding.EncodeToString(imageBytes)
	image := &model.Image{
		Data:      base64Image,
		UserId:    userId,
		EnityId:   entityId,
		ImageType: imageType,
	}
	err = s.repository.UpdateProfilePhoto(ctx, image)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) SaveImage(path string, imageType string, userId, entityId uuid.UUID) error {
	imageBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	base64Image := base64.StdEncoding.EncodeToString(imageBytes)
	image := &model.Image{
		Data:      base64Image,
		UserId:    userId,
		EnityId:   entityId,
		ImageType: imageType,
	}

	return s.repository.SaveImage(image)
}

func (s *service) GetProfilePhoto(ctx context.Context, userId uuid.UUID) (*model.Image, error) {
	image, err := s.repository.GetProfilePhoto(ctx, userId)
	if err != nil {
		return image, err
	}
	return image, nil
}
