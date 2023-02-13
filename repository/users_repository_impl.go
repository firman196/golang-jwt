package repository

import (
	"context"
	"database/sql"
	"golang-jwt/helper"
	"golang-jwt/model/entity"
)

type UsersRepositoryImpl struct{}

//func NewUsersRepositoryImpl() UsersRepository {
//	return &UsersRepository{}
//}

func (repository *UsersRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user entity.Users) entity.Users {
	SQL := "INSERT INTO users (id, first_name, last_name, email, password) VALUES (?,?,?,?,?)"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.Id,
		user.Firstname,
		user.Lastname,
		user.Email,
		user.Password,
	)
	helper.SetPanicError(err)
	return user
}
