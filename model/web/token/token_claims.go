package token

import "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	UserId    string `json:"user_id"`
	Email     string `json:"email"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	jwt.StandardClaims
}
