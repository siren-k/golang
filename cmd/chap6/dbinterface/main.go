package main

import (
	"golang/chap6/database"
	"golang/chap6/dbinterface"
)

func main() {
	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// 커밋을 설공하면 다음 코드는 아무 작업도 하지 않는다.
	defer tx.Rollback()

	if err = dbinterface.Exec(tx); err != nil {
		panic(err)
	}

	if err = tx.Commit(); err != nil {
		panic(err)
	}
}
