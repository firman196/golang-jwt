package helper

import (
	"database/sql"
	"golang-jwt/utils"
)

func Defer(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		utils.SetPanicError(errRollback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		utils.SetPanicError(errCommit)
	}
}
