package handler

import (
	"LayeredArchitecture/config"
	"LayeredArchitecture/interfaces/dddcontext"
	"LayeredArchitecture/interfaces/response"
	"LayeredArchitecture/usecase"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func HandlePostGet(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// パラメータからpostIDを取得
	postID, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
	}
	//applicationレイヤを操作して、ユーザデータ取得
	post, err := usecase.PostUsecase{}.SelectByPrimaryKey(config.DB, postID)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	response.JSON(writer, http.StatusOK, post)
}

func HandlePostsGet(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	//applicationレイヤを操作して、ユーザデータ取得
	posts, err := usecase.PostUsecase{}.GetAll(config.DB)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	response.JSON(writer, http.StatusOK, posts)
}

func HandlePostCreate(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	// Contextから認証済みのユーザIDを取得
	ctx := request.Context()
	userID := dddcontext.GetUserIDFromContext(ctx)

	//リクエストボディからサインアップ情報を取得
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err, "Invalid Request Body")
		return
	}

	//リクエストボディのパース
	var requestBody postRequest
	json.Unmarshal(body, &requestBody)
	err = usecase.PostUsecase{}.Insert(config.DB, requestBody.Content, userID)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	response.JSON(writer, http.StatusOK, "")
}

func HandlePostUpdate(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// パラメータからpostIDを取得
	postID, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
	}
	//リクエストボディからサインアップ情報を取得
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err, "Invalid Request Body")
		return
	}

	//リクエストボディのパース
	var requestBody postRequest
	json.Unmarshal(body, &requestBody)

	//applicationレイヤを操作して、ユーザデータ更新
	err = usecase.PostUsecase{}.UpdateByPrimaryKey(config.DB, postID, requestBody.Content)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	response.JSON(writer, http.StatusOK, "")
}

func HandlePostDelete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// パラメータからpostIDを取得
	postID, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
	}
	//applicationレイヤを操作して、ユーザデータ削除
	err = usecase.PostUsecase{}.DeleteByPrimaryKey(config.DB, postID)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	response.JSON(writer, http.StatusOK, "")
}

type postRequest struct {
	Content string `json:"content"`
}
