package dtos

import (
	"github.com/google/uuid"
)

type LoginReq struct {
	Identifier string `json:"identifier" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

type RenewPassReq struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
	UserID      uint   `json:"user_id"`
}

type ChangeForgotPassReq struct {
	Code       string `json:"code" validate:"required"`
	NewPass    string `json:"new_pass" validate:"required"`
	Identifier string `json:"identifier" validate:"required"`
}

type LoginDTO struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"username"`
	Token    string    `json:"token"`
}
