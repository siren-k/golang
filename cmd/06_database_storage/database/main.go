package main

import (
	database2 "golang/internal/06_database_storage/database"
)

func main() {
	db, err := database2.Setup()
	if err != nil {
		panic(err)
	}

	if err := database2.Exec(db); err != nil {
		panic(err)
	}
}
