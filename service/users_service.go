package service

import (
	"context"
	"golang-jwt/model/web/users"
)

type UsersService interface {
	Create(ctx context.Context, request users.UsersCreateRequest) users.UsersResponse
	Update(ctx context.Context, request users.UsersUpdateRequest) users.UsersResponse
	Delete(ctx context.Context, Id int16) bool
	GetById(ctx context.Context, Id int16) users.UsersResponse
	//GetByEmail(ctx context.Context, email string) users.UsersResponse
	GetAll(ctx context.Context) []users.UsersResponse
}
