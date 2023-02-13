package repository

import (
	"context"
	"database/sql"
	"golang-jwt/model/entity"
)

type UsersRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user entity.Users) entity.Users
	Update(ctx context.Context, tx *sql.Tx, user entity.Users) entity.Users
	Delete(ctx context.Context, tx *sql.Tx, user entity.Users) bool
	getById(ctx context.Context, tx *sql.Tx, userId string) (entity.Users, error)
	getByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.Users, error)
	getAll(ctx context.Context, tx *sql.Tx) []entity.Users
}
