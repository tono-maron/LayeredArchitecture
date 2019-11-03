package usecase

import (
	"LayeredArchitecture/domain"
	"LayeredArchitecture/domain/repository"
)

type PostUsecase interface {
	SelectByPrimaryKey(postID int) (*domain.Post, error)
	GetAll() ([]domain.Post, error)
	Insert(content, userID string) error
	UpdateByPrimaryKey(postID int, content string) error
	DeleteByPrimaryKey(postID int) error
}

type postUsecase struct {
	postRepository repository.PostRepository
}

// NewUserUseCase : User データに関する UseCase を生成
func NewPostUsecase(pr repository.PostRepository) PostUsecase {
	return &postUsecase{
		postRepository: pr,
	}
}

func (pu postUsecase) SelectByPrimaryKey(postID int) (*domain.Post, error) {
	post, err := pu.postRepository.SelectByPrimaryKey(postID)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (pu postUsecase) GetAll() ([]domain.Post, error) {
	posts, err := pu.postRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (pu postUsecase) Insert(content, userID string) error {
	err := pu.postRepository.Insert(content, userID)
	if err != nil {
		return err
	}
	return nil
}

func (pu postUsecase) UpdateByPrimaryKey(postID int, content string) error {
	err := pu.postRepository.UpdateByPrimaryKey(postID, content)
	if err != nil {
		return err
	}
	return nil
}

func (pu postUsecase) DeleteByPrimaryKey(postID int) error {
	err := pu.postRepository.DeleteByPrimaryKey(postID)
	if err != nil {
		return err
	}
	return nil
}
