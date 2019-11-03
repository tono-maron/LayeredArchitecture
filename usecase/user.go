package usecase

import (
	"LayeredArchitecture/domain"
	"LayeredArchitecture/domain/repository"
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// UserUseCase : User における UseCase のインターフェース
type UserUsecase interface {
	SelectByPrimaryKey(userID string) (*domain.User, error)
	Insert(name, email, password string) error
	CreateAuthToken(email, password string) (string, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

// NewUserUseCase : User データに関する UseCase を生成
func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (uu userUsecase) SelectByPrimaryKey(userID string) (*domain.User, error) {
	user, err := uu.userRepository.SelectByPrimaryKey(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu userUsecase) Insert(name, email, password string) error {
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

	err = uu.userRepository.Insert(userID.String(), name, email, string(passwordDigest), false)
	if err != nil {
		return err
	}
	return nil
}

func (uu userUsecase) CreateAuthToken(email, password string) (string, error) {
	user, err := uu.userRepository.SelectByEmail(email)
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
