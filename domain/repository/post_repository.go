package repository

import (
	"database/sql"

	"LayeredArchitecture/domain"
)

type PostRepository interface {
	func SelectByPrimaryKey(DB *sql.DB, postID int) (*domain.Post, error)
	func GetAll(DB *sql.DB) ([]domain.Post, error)
	func Insert(DB *sql.DB, postID int, content, userID string) error
	func UpdateByPrimaryKey(DB *sql.DB, postID int, content string) error
	func Delete(DB *sql.DB, postID int) error
}