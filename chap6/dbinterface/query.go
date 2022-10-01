package dbinterface

import (
	"fmt"
	"golang/chap6/database"
)

// Query 함수는 새로운 연결을 포착해 테이블을 생성하고
// 몇 가지 쿼리를 수행한 다음, 나중에 테이블을 제거한다.
func Query(db DB) error {
	name := "Aaron"
	rows, err := db.Query("SELECT name, created FROM example WHERE name=?", name)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var e database.Example
		if err = rows.Scan(&e.Name, &e.Created); err != nil {
			return err
		}
		fmt.Printf("Results:\n\tName: %s\n\tCreated: %v\n", e.Name, e.Created)
	}

	return rows.Err()
}
