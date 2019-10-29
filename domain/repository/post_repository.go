package repository

import (
	"database/sql"

	"LayeredArchitecture/domain"
)

type PostRepository interface {
	SelectByPrimaryKey(DB *sql.DB, postID int) (*domain.Post, error)
	GetAll(DB *sql.DB) ([]domain.Post, error)
	Insert(DB *sql.DB, content, userID string) error
	UpdateByPrimaryKey(DB *sql.DB, postID int, content string) error
	DeleteByPrimaryKey(DB *sql.DB, postID int) error
}
