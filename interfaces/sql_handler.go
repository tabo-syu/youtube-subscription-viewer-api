package interfaces

import "database/sql"

type SqlHandler interface {
	Exec(query string, args ...any) (sql.Result, error)
}
