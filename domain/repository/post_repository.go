package repository

import (
	"LayeredArchitecture/domain"
)

type PostRepository interface {
	SelectByPrimaryKey(postID int) (*domain.Post, error)
	GetAll() ([]domain.Post, error)
	Insert(content, userID string) error
	UpdateByPrimaryKey(postID int, content string) error
	DeleteByPrimaryKey(postID int) error
}
