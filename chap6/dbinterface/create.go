package dbinterface

import _ "github.com/go-sql-driver/mysql"

// Create 함수는 example이라는 이름의 테이블을 만들고
// 데이터를 추가한다.
func Create(db DB) error {
	// 데이터베이스 생성
	if _, err := db.Exec("CREATE TABLE example (name VARCHAR(20), created DATETIME)"); err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO example (name, created) values ("Aaron", NOW())`); err != nil {
		return err
	}

	return nil
}
