package pools

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Setup 함수는 pools 패키지를 사용하고 연결 수 제한 등을 적용해 db를 구성한다.
func Setup() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/chapter6_db?parseTime=true", "gopher", "gopher"))
	if err != nil {
		return nil, err
	}

	// 최대 24개의 연결만 허용한다.
	db.SetMaxOpenConns(24)

	// MaxIdleConns는 열려 있는 최대 SetMaxOpenConns보다 작을 수 없다.
	// 그렇지 않으면 해당 값을 기본 설정된다.
	db.SetMaxIdleConns(24)

	return db, nil
}
