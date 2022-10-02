package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Create 함수는 example이라는 이름의 테이블을 생성하고 데이터를 채운다.
func Create(db *sql.DB) error {
	// 데이터베이스 생성하기
	if _, err := db.Exec("CREATE TABLE example (name VARCHAR(20), created DATETIME)"); err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO example (name, created) values ("Aaron", NOW())`); err != nil {
		return err
	}

	return nil
}
