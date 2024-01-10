package model

import "social-media-app/pkg/utils"

type User struct {
	Base
	Username       string `gorm:"uniqueIndex;not null" json:"username"`
	Phone          string `gorm:"uniqueIndex;not null" json:"phone"`
	Email          string `gorm:"uniqueIndex;not null" json:"email"`
	Password       string `gorm:"not null" json:"password"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
}

func (u *User) PassHash() error {
	pass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = pass
	return nil
}
