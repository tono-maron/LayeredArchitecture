package handler

import (
	"LayeredArchitecture/interfaces/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	response.JSON(writer, http.StatusOK, "GO DDD API")
}
