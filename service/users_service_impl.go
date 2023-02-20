package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"

	"golang-jwt/helper"
	"golang-jwt/model/web/users"
	"golang-jwt/repository"
)

type UsersServiceImpl struct {
	UsersRepository repository.UsersRepository
	DB              *sql.DB
	validate        *validator.Validate
}

/*func NewUserServiceImpl(userRepository repository.UsersRepository, DB *sql.DB, validate *validator.Validate) UsersService {
	return &UsersServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		validate:       validate,
	}
}*/

func (service *UsersServiceImpl) Create(ctx context.Context, request users.UsersCreateRequest) users.UsersResponse {
	err := service.validate.Struct(request)

	helper.SetPanicError(err)
	tx, err := service.DB.Begin()
	helper.SetPanicError(err)
	defer helper.Defer(tx)

}
