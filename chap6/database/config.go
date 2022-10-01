package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// Example 구조체는 쿼리에 대한 결과를 저장한다.
type Example struct {
	Name    string
	Created *time.Time
}

// Setup 함수는 데이터베이스를 설정하고
// 연결 풀이 설정된 데이터베이스를 반환한다.
func Setup() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/chapter6_db?parseTime=true", "gopher", "gopher"))
	if err != nil {
		return nil, err
	}

	return db, nil
}
