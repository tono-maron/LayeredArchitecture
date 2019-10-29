package handler

import (
	"LayeredArchitecture/config"
	"LayeredArchitecture/interfaces/response"
	"LayeredArchitecture/usecase"
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

}

func HandlePostUpdate(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

}

func HandlePostDelete(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

}
