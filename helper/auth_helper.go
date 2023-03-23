package helper

import (
	"golang-jwt/model/web/token"
	"golang-jwt/model/web/users"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(request users.UsersResponse, value time.Duration) string {
	//var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))
	var APPLICATION_NAME = "BELAJAR"

	expiredTime := time.Now().Add(time.Minute * value).Unix()

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
	tokenStr, err := tokens.SignedString([]byte("dsgdcyusd6stsatts5rsf"))
	SetPanicError(err)

	return tokenStr

}
