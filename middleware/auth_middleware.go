package middleware

import (
	exception "golang-jwt/exception/api"
	"golang-jwt/model/web/token"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" && (request.RequestURI == "/api/v1/user" || request.RequestURI == "/api/v1/auth") {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		tokenAuth := request.Header.Get("Authorization")
		if tokenAuth == "" {
			panic(exception.NewUnautorizedRequestError("User Unautorized"))
		}

		var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))

		claims := &token.TokenClaims{}

		token, err := jwt.ParseWithClaims(tokenAuth, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtTokenSecret, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				panic(exception.NewUnautorizedRequestError(err.Error()))
			}
		}

		if !token.Valid {
			panic(exception.NewUnautorizedRequestError(err.Error()))
		}

		middleware.Handler.ServeHTTP(writer, request)
	}

}
