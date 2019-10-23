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
