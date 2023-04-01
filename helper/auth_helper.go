package helper

import (
	exception "golang-jwt/exception/api"
	"golang-jwt/model/web/token"
	"golang-jwt/model/web/users"
	"golang-jwt/utils"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(request users.UsersResponse, value time.Duration) string {
	var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))
	var APPLICATION_NAME = "BELAJAR"

	expiredTime := time.Now().Add(time.Hour * value).Unix()

	claims := &token.TokenClaims{
		Id:        request.Id,
		Email:     request.Email,
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: expiredTime,
		},
	}

	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := tokens.SignedString(jwtTokenSecret)
	if err != nil {
		utils.SetPanicError(err)
	}
	return tokenStr

}

func TokenClaims(userToken string) token.TokenClaims {
	var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))
	claims := &token.TokenClaims{}

	token, err := jwt.ParseWithClaims(userToken, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtTokenSecret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			panic(exception.NewUnauthorizedError(err.Error()))
		}
	}

	if !token.Valid {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	return *claims

}
