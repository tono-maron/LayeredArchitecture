package handler

import (
	"LayeredArchitecture/interfaces/dcontext"
	"LayeredArchitecture/interfaces/response"
	"LayeredArchitecture/usecase"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type PostHandler interface {
	HandlePostGet(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	HandlePostCreate(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	HandlePostsGet(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	HandlePostUpdate(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	HandlePostDelete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type postHandler struct {
	postUsecase usecase.PostUsecase
}

// NewUserUsecase : User データに関する Handler を生成
func NewPostHandler(pu usecase.PostUsecase) PostHandler {
	return &postHandler{
		postUsecase: pu,
	}
}

func (ph postHandler) HandlePostGet(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// パラメータからpostIDを取得
	postID, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
	}
	//applicationレイヤを操作して、ユーザデータ取得
	post, err := ph.postUsecase.SelectByPrimaryKey(postID)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	response.JSON(writer, http.StatusOK, post)
}

func (ph postHandler) HandlePostsGet(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	//applicationレイヤを操作して、ユーザデータ取得
	posts, err := ph.postUsecase.GetAll()
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	response.JSON(writer, http.StatusOK, posts)
}

func (ph postHandler) HandlePostCreate(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	// Contextから認証済みのユーザIDを取得
	userID := dcontext.GetUserIDFromContext(dcontext.Ctx)

	//リクエストボディからサインアップ情報を取得
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err, "Invalid Request Body")
		return
	}

	//リクエストボディのパース
	var requestBody postRequest
	json.Unmarshal(body, &requestBody)
	err = ph.postUsecase.Insert(requestBody.Content, userID)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	response.JSON(writer, http.StatusOK, "")
}

func (ph postHandler) HandlePostUpdate(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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
	err = ph.postUsecase.UpdateByPrimaryKey(postID, requestBody.Content)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	response.JSON(writer, http.StatusOK, "")
}

func (ph postHandler) HandlePostDelete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// パラメータからpostIDを取得
	postID, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
	}
	//applicationレイヤを操作して、ユーザデータ削除
	err = ph.postUsecase.DeleteByPrimaryKey(postID)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	response.JSON(writer, http.StatusOK, "")
}

type postRequest struct {
	Content string `json:"content"`
}
