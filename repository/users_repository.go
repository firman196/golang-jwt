package repository

import (
	"context"
	"database/sql"
	"golang-jwt/model/entity"
)

type UsersRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user entity.Users) entity.Users
	Update(ctx context.Context, tx *sql.Tx, userId int16, user entity.Users) entity.Users
	Delete(ctx context.Context, tx *sql.Tx, userId int16) bool
	GetById(ctx context.Context, tx *sql.Tx, userId int16) (entity.Users, error)
	GetByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.Users, error)
	GetAll(ctx context.Context, tx *sql.Tx) []entity.Users
}
