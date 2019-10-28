package handler

import (
	"LayeredArchitecture/config"
	"LayeredArchitecture/interfaces/response"
	"LayeredArchitecture/usecase"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HandleUserGet(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	user, err := usecase.UserUsecase{}.SelectByPrimaryKey(config.DB, userID)
	if err != nil {
		log.Println(err)
		response.Error(writer, http.StatusInternalServerError, "Internal Server Error")
	}
	response.JSON(writer, http.StatusOK, user)
}
