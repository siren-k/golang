package dbinterface

import "database/sql"

// DB는 sql.DB 또는 sql.Transaction을 만족하는 인터페이스다.
type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// Transaction 인터페이스는 커밋(Commit), 되돌리기(Rollback), Stmt 등
// 쿼리(Query)가 할 수 있는 모든 작업을 할 수 있다.
type Transaction interface {
	DB
	Commit() error
	Rollback() error
}
