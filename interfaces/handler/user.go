package handler

import (
	"LayeredArchitecture/config"
	"LayeredArchitecture/interfaces/dddcontext"
	"LayeredArchitecture/interfaces/response"
	"LayeredArchitecture/usecase"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HandleUserGet(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	// Contextから認証済みのユーザIDを取得
	ctx := request.Context()
	//userIDが空かどうかのチェックはミドルウェアで実装してあるためここでのエラーハンドリングはない。
	userID := dddcontext.GetUserIDFromContext(ctx)

	user, err := usecase.UserUsecase{}.SelectByPrimaryKey(config.DB, userID)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	response.JSON(writer, http.StatusOK, user)
}
