package handler

import (
	"LayeredArchitecture/config"
	"LayeredArchitecture/domain"
	"LayeredArchitecture/interfaces/dddcontext"
	"LayeredArchitecture/interfaces/response"
	"LayeredArchitecture/usecase"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserHandler interface {
	SelectByPrimaryKey(DB *sql.DB, userID string) (*domain.User, error)
	Insert(DB *sql.DB, userID, name, email, password string, admin bool) error
	SelectByEmail(DB *sql.DB, email string) (*domain.User, error)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

// NewUserUsecase : User データに関する Handler を生成
func NewBookHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: uu,
	}
}

//ユーザ情報取得
func (uh userHandler) HandleUserGet(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	// Contextから認証済みのユーザIDを取得
	ctx := request.Context()
	userID := dddcontext.GetUserIDFromContext(ctx)

	//applicationレイヤを操作して、ユーザデータ取得
	user, err := usecase.UserUsecase{}.SelectByPrimaryKey(config.DB, userID)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	response.JSON(writer, http.StatusOK, user)
}

// "/user/signup" 新規登録
func (uh userHandler) HandleUserSignup(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	//リクエストボディからサインアップ情報を取得
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err, "Invalid Request Body")
		return
	}

	//リクエストボディのパース
	var requestBody userSignupRequest
	json.Unmarshal(body, &requestBody)

	//userIDによってuserテーブルにハッシュ化されたパスワードとemaiと更新されたauth_tokenを更新する
	err = usecase.UserUsecase{}.Insert(config.DB, requestBody.Name, requestBody.Email, requestBody.Password)
	if err != nil {
		log.Println(err)
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	// レスポンスに必要な情報を詰めて返却
	response.JSON(writer, http.StatusOK, "")
}

//"/user/signin" ログイン機能
func (uh userHandler) HandleUserSignin(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	// リクエストBodyからログイン情報を取得
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err, "Invalid Request Body")
		return
	}
	//リクエストボディのパース
	var requestBody userLoginRequest
	json.Unmarshal(body, &requestBody)

	//Emailによってユーザ情報取得し、そこから認証トークンを作成し取得する。
	authToken, err := usecase.UserUsecase{}.CreateAuthToken(config.DB, requestBody.Email, requestBody.Password)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}

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
