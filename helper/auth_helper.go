package helper

import (
	"golang-jwt/model/web/token"
	"golang-jwt/model/web/users"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(request users.UsersResponse, value time.Duration) string {
	var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))

	expiredTime := time.Now().Add(time.Minute * value)
	claims := &token.TokenClaims{
		Id:        request.Id,
		Email:     request.Email,
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenStr, err := token.SignedString(jwtTokenSecret)
	SetPanicError(err)

	return tokenStr

}
