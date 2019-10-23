package repository

import (
	"database/sql"

	"LayeredArchitecture/domain"
)

type UserRepository interface {
	SelectByPrimaryKey(DB *sql.DB, userID string) (*domain.User, error)
}
