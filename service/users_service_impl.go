package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"

	"golang-jwt/helper"
	"golang-jwt/model/entity"
	"golang-jwt/model/web/users"
	"golang-jwt/repository"
)

type UsersServiceImpl struct {
	UsersRepository repository.UsersRepository
	DB              *sql.DB
	validate        *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UsersRepository, DB *sql.DB, validate *validator.Validate) UsersService {
	return &UsersServiceImpl{
		UsersRepository: userRepository,
		DB:              DB,
		validate:        validate,
	}
}

func (service *UsersServiceImpl) Create(ctx context.Context, request users.UsersCreateRequest) users.UsersResponse {
	err := service.validate.Struct(request)
	helper.SetPanicError(err)

	tx, err := service.DB.Begin()
	helper.SetPanicError(err)
	defer helper.Defer(tx)

	passwordHash, err := helper.HashPassword(request.Password)
	helper.SetPanicError(err)

	user := entity.Users{
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Email:     request.Email,
		Password:  passwordHash,
	}

	user = service.UsersRepository.Create(
		ctx,
		tx,
		user,
	)

	return users.UsersResponses(user)
}

func (service *UsersServiceImpl) Update(ctx context.Context, request users.UsersUpdateRequest) users.UsersResponse {
	err := service.validate.Struct(request)
	helper.SetPanicError(err)

	tx, err := service.DB.Begin()
	helper.SetPanicError(err)
	defer helper.Defer(tx)

	user, err := service.UsersRepository.GetById(
		ctx,
		tx,
		request.Id,
	)
	if err != nil {
		helper.SetPanicError(err)
	}
	user.Firstname = request.Firstname
	user.Lastname = request.Lastname
	user.Email = request.Email

	user = service.UsersRepository.Update(
		ctx,
		tx,
		user,
	)

	return users.UsersResponses(user)
}

func (service *UsersServiceImpl) Delete(ctx context.Context, Id int16) bool {
	tx, err := service.DB.Begin()
	helper.SetPanicError(err)
	defer helper.Defer(tx)

	user, err := service.UsersRepository.GetById(
		ctx,
		tx,
		Id,
	)
	helper.SetPanicError(err)

	var result = service.UsersRepository.Delete(
		ctx,
		tx,
		user.Id,
	)

	return result
}

func (service *UsersServiceImpl) GetById(ctx context.Context, Id int16) users.UsersResponse {
	tx, err := service.DB.Begin()
	helper.SetPanicError(err)
	defer helper.Defer(tx)

	user, err := service.UsersRepository.GetById(
		ctx,
		tx,
		Id,
	)
	helper.SetPanicError(err)
	return users.UsersResponses(user)
}

func (service *UsersServiceImpl) GetAll(ctx context.Context) []users.UsersResponse {
	tx, err := service.DB.Begin()
	helper.SetPanicError(err)
	defer helper.Defer(tx)

	usersVal := service.UsersRepository.GetAll(
		ctx,
		tx,
	)
	var usersResponse []users.UsersResponse
	for _, user := range usersVal {
		usersResponse = append(usersResponse, users.UsersResponses(user))
	}
	return usersResponse
}
