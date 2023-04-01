package service

import (
	"context"
	"golang-jwt/model/web/token"
	"golang-jwt/model/web/users"
)

type UsersService interface {
	Create(ctx context.Context, request users.UsersCreateRequest) users.UsersResponse
	Update(ctx context.Context, Id int16, request users.UsersUpdateRequest) users.UsersResponse
	Delete(ctx context.Context, Id int16) bool
	GetById(ctx context.Context, Id int16) users.UsersResponse
	Auth(ctx context.Context, request users.UserAuthRequest) token.TokenResponse
	GetAll(ctx context.Context) []users.UsersResponse
	RefreshToken(ctx context.Context, refreshToken string) token.TokenResponse
}
