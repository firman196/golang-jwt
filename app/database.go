package app

import (
	"database/sql"
	"golang-jwt/helper"
	"time"
)

func Database() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/yt_users_service")
	helper.SetPanicError(err)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
