package usecase

import (
	"LayeredArchitecture/domain"
	"LayeredArchitecture/domain/repository"
	"LayeredArchitecture/infrastructure/persistence"
	"database/sql"
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	passwordDigest, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	//UUIDでユーザIDを取得
	userID, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	err = repository.UserRepository(persistence.UserPersistence{}).Insert(DB, userID.String(), name, email, string(passwordDigest), false)
	if err != nil {
		return err
	}
	return nil
}

func (userUsecase UserUsecase) CreateAuthToken(DB *sql.DB, email, password string) (string, error) {
	user, err := repository.UserRepository(persistence.UserPersistence{}).SelectByEmail(DB, email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}
	authToken, err := createJWT(user.UserID)
	if err != nil {
		return "", err
	}
	return authToken, nil
}

//JWTで利用する署名
var Signature = "lt9m2-vn8bzf-02p-sgaq-32r9hdvanva"

func createJWT(userID string) (string, error) {
	//認証トークンを生成する
	//UUIDでユーザIDを取得
	internalToken, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)
	//claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = userID
	claims["it"] = internalToken
	// 電子署名
	tokenString, err := token.SignedString([]byte(Signature))
	if err != nil {
		return "", err
	}
	strByte := []byte(tokenString)
	authToken := string(strByte)
	return authToken, nil
}
