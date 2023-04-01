package token

import "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	Id        int16  `json:"id"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	jwt.StandardClaims
}
