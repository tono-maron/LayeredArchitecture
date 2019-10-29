package handler

import (
	"LayeredArchitecture/config"
	"LayeredArchitecture/interfaces/dddcontext"
	"LayeredArchitecture/interfaces/middleware"
	"LayeredArchitecture/interfaces/response"
	"LayeredArchitecture/usecase"
	"encoding/json"
	"io/ioutil"
	"net/http"
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



// "/user/signup" 新規登録
func HandleUserSignup(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
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
	err = usecase.UserUsecase{}.Insert(config.DB,　requestBody.Name, requestBody.Email, requestBody.Password)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	// レスポンスに必要な情報を詰めて返却
	response.JSON(writer, http.StatusOK, "")
}



//"/user/signin" ログイン機能
func HandleUserSignin(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
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
