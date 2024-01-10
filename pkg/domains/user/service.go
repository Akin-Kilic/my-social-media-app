package user

import (
	"context"
	"errors"
	"fmt"
	"social-media-app/pkg/config"
	"social-media-app/pkg/constant"
	"social-media-app/pkg/dtos"
	"social-media-app/pkg/model"
	"social-media-app/pkg/redis"
	"social-media-app/pkg/utils"
	"time"
)

type Service interface {
	Register(ctx context.Context, user *model.User) error
	Login(ctx context.Context, req dtos.LoginReq) (dtos.LoginDTO, error)
	//Logout(ctx context.Context, user *model.User) error
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

	err = s.repository.Save(ctx, user)
	if err != nil {
		return errors.New(constant.FailedCreateUser)
	}
	return nil
}
func (s *service) Login(ctx context.Context, req dtos.LoginReq) (dtos.LoginDTO, error) {
	var loginDto dtos.LoginDTO
	user, err := s.repository.FindUserWithIdentifiers(ctx, req.Identifier)
	if err != nil {
		return loginDto, errors.New(constant.FailedLogin)
	}

	isTrue := utils.PasswordControl(user.Password, req.Password)
	if !isTrue {
		return loginDto, errors.New(constant.FailedUserNameOrPass)

	}
	ip := ctx.Value("ip_address").(string)
	token, err := utils.GenerateJwt(user.ID, config.ReadValue().JwtSecret, ip)
	if err != nil {
		return loginDto, errors.New(constant.FailedLogin)
	}
	//rds := redis.Client()

	key := fmt.Sprintf(constant.RedisForJwt, token, user.ID)
	redis.Set(key, token, time.Hour*time.Duration(config.ReadValue().JwtExpTime))

	loginDto.Convert(&user, token)

	return loginDto, nil

}

func (s *service) Logout(ctx context.Context, user *model.User) error {
	
	

	err = s.repository.Save(ctx, user)
	if err != nil {
		return errors.New(constant.FailedCreateUser)
	}
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

// func createUser(s service) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		var (
// 			payload dtos.UserReqDto
// 			resp    dtos.UserRespDto
// 		)
// 		if err := c.BodyParser(&payload); err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload" + err.Error()})
// 		}
// 		errs := utils.ValidateStruct(&payload)
// 		if errs != nil {
// 			return c.Status(400).JSON(utils.Response(errs))
// 		}
// 		err := s.CreateUser(&payload, &resp)
// 		if err != nil {
// 			return c.Status(404).JSON(utils.Response(err.Error()))
// 		}
// 		resp.Message = "User created successfully"
// 		return c.Status(200).JSON(resp)
// 	}
// }
