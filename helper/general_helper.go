package helper

import "database/sql"

func SetPanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func Defer(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		SetPanicError(errRollback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		SetPanicError(errCommit)
	}
}
