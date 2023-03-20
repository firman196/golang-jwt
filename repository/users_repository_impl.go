package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-jwt/helper"
	"golang-jwt/model/entity"
)

type UsersRepositoryImpl struct{}

func NewUsersRepositoryImpl() UsersRepository {
	return &UsersRepositoryImpl{}
}

func (repository *UsersRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user entity.Users) entity.Users {
	SQL := "INSERT INTO users (id, firstname, lastname, email, password) VALUES (?,?,?,?,?)"
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

func (repository *UsersRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, userId int16, user entity.Users) entity.Users {
	SQL := "UPDATE users SET firstname = ?, lastname = ?, email = ? WHERE id = ?"

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

func (repository *UsersRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, userId int16) (entity.Users, error) {
	SQL := "SELECT id, firstname, lastname, email FROM users WHERE id = ?"
	rows, err := tx.QueryContext(
		ctx,
		SQL,
		userId,
	)

	helper.SetPanicError(err)
	defer rows.Close()

	user := entity.Users{}
	if rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Firstname,
			&user.Lastname,
			&user.Email,
		)

		helper.SetPanicError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (repository UsersRepositoryImpl) GetByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.Users, error) {
	SQL := "SELECT id, first_name, last_name, email FROM users WHERE email = ?"

	rows, err := tx.QueryContext(
		ctx,
		SQL,
		email,
	)

	helper.SetPanicError(err)
	defer rows.Close()

	user := entity.Users{}
	if rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Firstname,
			&user.Lastname,
			&user.Email,
		)
		helper.SetPanicError(err)

		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (repository *UsersRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []entity.Users {
	SQL := "SELECT id, firstname, lastname, email FROM users"
	rows, err := tx.QueryContext(
		ctx,
		SQL,
	)

	helper.SetPanicError(err)
	defer rows.Close()

	var users []entity.Users
	for rows.Next() {
		user := entity.Users{}
		err := rows.Scan(
			&user.Id,
			&user.Firstname,
			&user.Lastname,
			&user.Email,
		)
		helper.SetPanicError(err)

		users = append(users, user)
	}

	return users

}
