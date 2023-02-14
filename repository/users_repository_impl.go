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

func (repository *UsersRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user entity.Users) entity.Users {
	SQL := "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?"

	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.Firstname,
		user.Lastname,
		user.Email,
		user.Id,
	)

	helper.SetPanicError(err)
	return user
}

func (repository *UsersRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userId int16) bool {
	SQL := "DELETE FROM users WHERE id = ?"

	_, err := tx.ExecContext(
		ctx,
		SQL,
		userId,
	)

	helper.SetPanicError(err)

	return true
}

func (repository *UsersRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, userId int16) entity.Users {
	SQL := "SELECT id, first_name, last_name, email FROM users WHERE id = ?"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		userId,
	)

	helper.SetPanicError(err)
	return entity.Users{}
}
