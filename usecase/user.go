package usecase

import (
	"LayeredArchitecture/domain"
	"LayeredArchitecture/domain/repository"
	"LayeredArchitecture/infrastructure/persistence"
	"database/sql"
)

type UserUsecase struct{}

func (userUsecase UserUsecase) SelectByPrimaryKey(DB *sql.DB, userID string) (*domain.User, error) {
	user, err := repository.UserRepository(persistence.UserPersistence{}).SelectByPrimaryKey(DB, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userUsecase UserUsecase) Insert(DB *sql.DB, userID, name, email, pass string, admin bool) error {
	err := repository.UserRepository(persistence.UserPersistence{}).Insert(DB, userID, name, email, pass, admin)
	if err != nil {
		return err
	}
	return nil
}

func (userUsecase UserUsecase) SelectByEmail(DB *sql.DB, email string) (*domain.User, error) {
	user, err := repository.UserRepository(persistence.UserPersistence{}).SelectByEmail(DB, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
