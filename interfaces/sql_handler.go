package interfaces

import "database/sql"

type SQLHandler interface {
	Exec(string, ...any) (sql.Result, error)
	QueryRow(string, ...any) *sql.Row
	Begin() (*sql.Tx, error)
}
