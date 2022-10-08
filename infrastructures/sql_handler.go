package infrastructures

import (
	"database/sql"
	"fmt"

	// PostgreSQL のドライバーをインポート。
	_ "github.com/jackc/pgx/v5/stdlib"
)

type SQLHandler struct {
	db *sql.DB
}

func NewSQLHandler(config *DB) (*SQLHandler, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable Timezone=%s",
		config.Host,
		config.User,
		config.Password,
		config.Name,
		config.Port,
		config.TimeZone,
	)

	postgres, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := postgres.Ping(); err != nil {
		return nil, err
	}

	return &SQLHandler{postgres}, nil
}

func (s *SQLHandler) Close() error {
	return s.db.Close()
}

func (s *SQLHandler) Exec(query string, args ...any) (sql.Result, error) {
	return s.db.Exec(query, args...)
}

func (s *SQLHandler) QueryRow(query string, args ...any) *sql.Row {
	return s.db.QueryRow(query, args...)
}

func (s *SQLHandler) Begin() (*sql.Tx, error) {
	return s.db.Begin()
}
