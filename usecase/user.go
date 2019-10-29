package usecase

import (
	"LayeredArchitecture/domain"
	"LayeredArchitecture/domain/repository"
	"LayeredArchitecture/infrastructure/persistence"
	"LayeredArchitecture/interfaces/response"
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct{}

func (userUsecase UserUsecase) SelectByPrimaryKey(DB *sql.DB, userID string) (*domain.User, error) {
	user, err := repository.UserRepository(persistence.UserPersistence{}).SelectByPrimaryKey(DB, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userUsecase UserUsecase) Insert(DB *sql.DB, name, email, password string) error {
	//passwordとemailのバリデーション
	if len(password) < 8 {
		return errors.New("validation error for password")
	}
	//TODO: しっかりバリデーションをする
	if !(strings.Contains(email, "@")) {
		return errors.New("validation error for email")
	}
	//パスワードをハッシュ化する
	var passwordDigest []byte
	passwordDigest, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	//UUIDでユーザIDを取得
	userID, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	err := repository.UserRepository(persistence.UserPersistence{}).Insert(DB, userID.string(), name, email, string(passwordDigest), admin)
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
