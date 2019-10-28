package middleware

import (
	"LayeredArchitecture/interfaces/dddcontext"
	"LayeredArchitecture/interfaces/response"
	"context"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

type JWT struct {
	Admin bool   `json:"admin"`
	Sub   string `json:"sub"`
	Name  string `json:"name"`
	Iat   string `json:"iat"`
	Exp   string `json:"exp"`
}

//改ざん検知のシグニチャ
const Signature string = "767f4534-v840-b523220x4a24"

func Authenticate(nextFunc httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// Get the Basic Authentication credentials
		//JWTを用いた認証
		ctx := request.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		// リクエストヘッダからx-token(認証トークン)を取得
		authToken := request.Header.Get("x-token")
		if authToken == "" {
			writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			response.Error(writer, http.StatusBadRequest, errors.New("x-token is empty"), "Bad Request")
			return
		}

		//JWTの解析で改ざんされていないか、有効期限が切れていないかを確認
		//利用されているアルゴリズムを確認
		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				err := errors.New("Unexpected signing method")
				writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
				response.Error(writer, http.StatusBadRequest, err, "Bad Request")
				return nil, err
			}
			return []byte(Signature), nil
		})
		if err != nil {
			writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			response.Error(writer, http.StatusBadRequest, errors.New("invalid token"), "Bad Request")
			return
		}

		//トークンの有効期限を確認
		if !token.Valid {
			writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			response.Error(writer, http.StatusBadRequest, errors.New("token is expired"), "Bad Request")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			response.Error(writer, http.StatusBadRequest, errors.New("cannot get claims"), "Bad Request")
			return
		}

		// 認証トークンからユーザIDを取得
		userID := claims["sub"].(string)
		if userID == "" {
			writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			response.Error(writer, http.StatusBadRequest, errors.New("userID is empty"), "Bad Request")
			return
		}

		// userIdをContextへ保存して以降の処理に利用する
		ctx = dddcontext.SetUserID(ctx, userID)

		//次の処理へ
		nextFunc(writer, request, params)
	}
}
