package handler

import (
	"LayeredArchitecture/config"
	"LayeredArchitecture/interfaces/dddcontext"
	"LayeredArchitecture/interfaces/middleware"
	"LayeredArchitecture/interfaces/response"
	"LayeredArchitecture/usecase"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

//ユーザ情報取得
func HandleUserGet(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	// Contextから認証済みのユーザIDを取得
	ctx := request.Context()
	//userIDが空かどうかのチェックはミドルウェアで実装してあるためここでのエラーハンドリングはない。
	userID := dddcontext.GetUserIDFromContext(ctx)

	//applicationレイヤを操作して、ユーザデータ取得
	user, err := usecase.UserUsecase{}.SelectByPrimaryKey(config.DB, userID)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	response.JSON(writer, http.StatusOK, user)
}

//新規登録
func HandleUserSignup(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	//リクエストBodyから更新後情報を取得
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err, "Invalid Request Body")
		return
	}

	//リクエストボディのパース
	var requestBody userSignupRequest
	json.Unmarshal(body, &requestBody)
	password := requestBody.Password
	email := requestBody.Email
	//passwordとemailのバリデーション
	if len(password) < 8 {
		response.Error(writer, http.StatusBadRequest, errors.New("validation error for password"), "Bad Request")
		return
	}
	//TODO: しっかりバリデーションをする
	if !(strings.Contains(email, "@")) {
		response.Error(writer, http.StatusBadRequest, errors.New("validation error for email"), "Bad Request")
		return
	}
	//パスワードをハッシュ化する
	var passwordDigest []byte
	passwordDigest, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}

	//UUIDでユーザIDを取得
	userID, err := uuid.NewRandom()
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}

	//userIDによってuserテーブルにハッシュ化されたパスワードとemaiと更新されたauth_tokenを更新する
	err = usecase.UserUsecase{}.Insert(config.DB, userID.String(), requestBody.Name, email, string(passwordDigest), false)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	// レスポンスに必要な情報を詰めて返却
	response.JSON(writer, http.StatusOK, "")
}

//ログイン
func HandleUserSignin(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	// リクエストBodyからログイン情報を取得
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err, "Invalid Request Body")
		return
	}
	var requestBody userLoginRequest
	json.Unmarshal(body, &requestBody)
	password := requestBody.Password
	email := requestBody.Email
	//一致するアカウント情報を取得
	user, err := usecase.UserUsecase{}.SelectByEmail(config.DB, email)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	//認証トークンを生成する
	//UUIDでユーザIDを取得
	internalToken, err := uuid.NewRandom()
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)
	//claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = user.UserID
	claims["it"] = internalToken
	// 電子署名
	tokenString, err := token.SignedString([]byte(middleware.Signature))
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	strByte := []byte(tokenString)
	authToken := string(strByte)

	// レスポンスに必要な情報を詰めて返却
	response.JSON(writer, http.StatusOK, tokenResponse{Token: authToken})
}

type userSignupRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type userLoginRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type tokenResponse struct {
	Token string `json:"token"`
}
