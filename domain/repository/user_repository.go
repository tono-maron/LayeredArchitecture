package repository

import (
	"LayeredArchitecture/domain"
)

type UserRepository interface {
	SelectByPrimaryKey(userID string) (*domain.User, error)
	Insert(userID, name, email, password string, admin bool) error
	SelectByEmail(email string) (*domain.User, error)
}
