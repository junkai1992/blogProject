package dto

import (
	"blogBack/model"
	"time"
)

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
