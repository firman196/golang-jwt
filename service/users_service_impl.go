package service

import (
	"context"
	"database/sql"
	"fmt"

	exception "golang-jwt/exception/api"
	"golang-jwt/helper"
	"golang-jwt/model/entity"
	"golang-jwt/model/web/users"
	"golang-jwt/repository"
)

type UsersServiceImpl struct {
	UsersRepository repository.UsersRepository
	DB              *sql.DB
	validate        *exception.Validation
}

func NewUserServiceImpl(userRepository repository.UsersRepository, DB *sql.DB, validate *exception.Validation) UsersService {
	return &UsersServiceImpl{
		UsersRepository: userRepository,
		DB:              DB,
		validate:        validate,
	}
}

func (service *UsersServiceImpl) Create(ctx context.Context, request users.UsersCreateRequest) users.UsersResponse {
	errValidation := service.validate.Struct(request)
	if errValidation != nil {
		panic(exception.NewBadRequestError(errValidation))
	}
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

func (service *UsersServiceImpl) Update(ctx context.Context, Id int16, request users.UsersUpdateRequest) users.UsersResponse {
	errValidation := service.validate.Struct(request)
	if errValidation != nil {
		panic(exception.NewBadRequestError(fmt.Sprint(errValidation)))
	}

	tx, err := service.DB.Begin()
	helper.SetPanicError(err)
	defer helper.Defer(tx)

	user, err := service.UsersRepository.GetById(
		ctx,
		tx,
		Id,
	)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}
	user.Firstname = request.Firstname
	user.Lastname = request.Lastname
	user.Email = request.Email

	user = service.UsersRepository.Update(
		ctx,
		tx,
		Id,
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

	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

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
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}
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
