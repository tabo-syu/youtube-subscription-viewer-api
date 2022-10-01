package interfaces

import "database/sql"

type SqlHandler interface {
	Exec(string, ...any) (sql.Result, error)
	QueryRow(string, ...any) *sql.Row
}
