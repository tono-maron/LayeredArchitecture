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
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type PostHandler interface {
	SelectByPrimaryKey(DB *sql.DB, postID int) (*domain.Post, error)
	GetAll(DB *sql.DB) ([]domain.Post, error)
	Insert(DB *sql.DB, content, userID string) error
	UpdateByPrimaryKey(DB *sql.DB, postID int, content string) error
	DeleteByPrimaryKey(DB *sql.DB, postID int) error
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
	post, err := usecase.PostUsecase{}.SelectByPrimaryKey(config.DB, postID)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	response.JSON(writer, http.StatusOK, post)
}

func (ph postHandler) HandlePostsGet(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	//applicationレイヤを操作して、ユーザデータ取得
	posts, err := usecase.PostUsecase{}.GetAll(config.DB)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	response.JSON(writer, http.StatusOK, posts)
}

func (ph postHandler) HandlePostCreate(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
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
