package repository

import (
	"database/sql"

	"LayeredArchitecture/domain"
)

type UserRepository interface {
	SelectByPrimaryKey(DB *sql.DB, userID string) (*domain.User, error)
	Insert(DB *sql.DB, userID, name, email, password string, admin bool) error
	SelectByEmail(DB *sql.DB, email string) (*domain.User, error)
}
