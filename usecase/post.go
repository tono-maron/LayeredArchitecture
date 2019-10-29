package usecase

import (
	"LayeredArchitecture/domain"
	"LayeredArchitecture/domain/repository"
	"LayeredArchitecture/infrastructure/persistence"
	"database/sql"
)

type PostUsecase struct{}

func (postUsecase PostUsecase) SelectByPrimaryKey(DB *sql.DB, postID int) (*domain.Post, error) {
	post, err := repository.PostRepository(persistence.PostPersistence{}).SelectByPrimaryKey(DB, postID)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (postUsecase PostUsecase) GetAll(DB *sql.DB) ([]domain.Post, error) {
	posts, err := repository.PostRepository(persistence.PostPersistence{}).GetAll(DB)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (postUsecase PostUsecase) Insert(DB *sql.DB, content, userID string) error {
	err := repository.PostRepository(persistence.PostPersistence{}).Insert(DB, content, userID)
	if err != nil {
		return err
	}
	return nil
}

func (postUsecase PostUsecase) UpdateByPrimaryKey(DB *sql.DB, postID int, content string) error {
	err := repository.PostRepository(persistence.PostPersistence{}).UpdateByPrimaryKey(DB, postID, content)
	if err != nil {
		return err
	}
	return nil
}

func (postUsecase PostUsecase) DeleteByPrimaryKey(DB *sql.DB, postID int) error {
	err := repository.PostRepository(persistence.PostPersistence{}).DeleteByPrimaryKey(DB, postID)
	if err != nil {
		return err
	}
	return nil
}
