package usecase

import (
	"LayeredArchitecture/domain"
	"LayeredArchitecture/domain/repository"
	"database/sql"
)

type PostUsecase interface {
	SelectByPrimaryKey(DB *sql.DB, postID int) (*domain.Post, error)
	GetAll(DB *sql.DB) ([]domain.Post, error)
	Insert(DB *sql.DB, content, userID string) error
	UpdateByPrimaryKey(DB *sql.DB, postID int, content string) error
	DeleteByPrimaryKey(DB *sql.DB, postID int) error
}

type postUsecase struct {
	postRepository repository.PostRepository
}

// NewUserUseCase : User データに関する UseCase を生成
func NewPostUseCase(pr repository.PostRepository) PostUsecase {
	return &postUsecase{
		postRepository: pr,
	}
}

func (pu postUsecase) SelectByPrimaryKey(DB *sql.DB, postID int) (*domain.Post, error) {
	post, err := pu.postRepository.SelectByPrimaryKey(DB, postID)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (pu postUsecase) GetAll(DB *sql.DB) ([]domain.Post, error) {
	posts, err := pu.postRepository.GetAll(DB)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (pu postUsecase) Insert(DB *sql.DB, content, userID string) error {
	err := pu.postRepository.Insert(DB, content, userID)
	if err != nil {
		return err
	}
	return nil
}

func (pu postUsecase) UpdateByPrimaryKey(DB *sql.DB, postID int, content string) error {
	err := pu.postRepository.UpdateByPrimaryKey(DB, postID, content)
	if err != nil {
		return err
	}
	return nil
}

func (pu postUsecase) DeleteByPrimaryKey(DB *sql.DB, postID int) error {
	err := pu.postRepository.DeleteByPrimaryKey(DB, postID)
	if err != nil {
		return err
	}
	return nil
}
