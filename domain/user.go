package domain

import (
	"infrastructure"
)

type User struct {
	UserID   int
	Name     string
	Email    string
	Password string
	Admin    bool
}

func (u *User) IsAdmin() bool {
   return u.Admin
}

func NewUserFromDTO(dto *UserDTO) *domain.User {
	user := &User{
		UserID:   dto.UserID,
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
		Admin:    dto.Admin,
	}
	return user
}