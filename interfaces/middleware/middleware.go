package middleware

import (
	"LayeredArchitecture/interfaces/dddcontext"
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
		user, authToken, hasAuth := request.BasicAuth()

		if hasAuth && user == requiredUser && password == requiredPassword {
			ctx := request.Context()
			if ctx == nil {
				ctx = context.Background()
			}

			// リクエストヘッダからx-token(認証トークン)を取得
			authToken := request.Header.Get("x-token")
			if authToken == "" {
				return "", errors.New("x-token is empty")
			}

			//JWTの解析で改ざんされていないか、有効期限が切れていないかを確認
			//利用されているアルゴリズムを確認
			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					err := errors.New("Unexpected signing method")
					return nil, err
				}
				return []byte(Signature), nil
			})
			if err != nil {
				return "", errors.New("invalid token")
			}

			//トークンの有効期限を確認
			if !token.Valid {
				return "", errors.New("Token is invalid")
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return "", errors.New("cannot get claims")
			}

			// 認証トークンからユーザIDを取得
			userID := claims["sub"].(string)
			if userID == "" {
				return "", errors.New("userID is empty")
			}

			// userIdをContextへ保存して以降の処理に利用する
			ctx = dddcontext.SetUserID(ctx, userID)

			return userID, nil
			// Delegate request to the given handle
			nextFunc(writer, request, params)
		} else {
			// Request Basic Authentication otherwise
			writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}
