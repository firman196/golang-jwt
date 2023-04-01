package app

import (
	"database/sql"
	"golang-jwt/utils"
	"time"
)

func Database() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang-jwt")
	utils.SetPanicError(err)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
