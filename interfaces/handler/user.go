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
	userID := "123"
	user, err := usecase.UserUsecase{}.SelectByPrimaryKey(config.DB, userID)
	if err != nil {
		log.Println(err)
	}
	response.JSON(writer, http.StatusOK, user)
}
