package usecase

import (
	"LayeredArchitecture/domain"
	"LayeredArchitecture/domain/repository"
	"LayeredArchitecture/infrastructure/persistence"
	"database/sql"
)

type PostUsecase struct{}

func (postUsecase PostUsecase) SelectByPrimaryKey(DB *sql.DB, postID int, userID string) (*domain.Post, error) {
	post, err := repository.PostRepository(persistence.PostPersistence{}).SelectByPrimaryKey(DB, postID, userID)
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

func (postUsecase PostUsecase) Insert(DB *sql.DB, content string) error {
	err := repository.PostRepository(persistence.PostPersistence{}).Insert(DB, content)
	if err != nil {
		return err
	}
	return nil
}

func (postUsecase PostUsecase) UpdateByPrimaryKey(DB *sql.DB, postID int, content string) error {

}

func (postUsecase PostUsecase) DeleteByPrimaryKey(DB *sql.DB, postID int) error {

}
