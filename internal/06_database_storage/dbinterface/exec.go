package dbinterface

// Exec 함수는 이전 예제의 Exec 함수를 대체한다.
func Exec(db DB) error {
	// 정리할 때 감지하지 않은 오류가 발생할 수 있다.
	// 하지만 항상 정리를 시도한다.
	defer db.Exec("DROP TABLE example")

	if err := Create(db); err != nil {
		return err
	}

	if err := Query(db); err != nil {
		return err
	}

	return nil
}
