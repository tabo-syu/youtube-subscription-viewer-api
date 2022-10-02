package infrastructures

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type SqlHandler struct {
	db *sql.DB
}

func NewSqlHandler(config *DB) (*SqlHandler, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable Timezone=%s",
		config.Host,
		config.User,
		config.Password,
		config.Name,
		config.Port,
		config.TimeZone,
	)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &SqlHandler{db}, nil
}

func (s *SqlHandler) Close() error {
	return s.db.Close()
}

func (s *SqlHandler) Exec(query string, args ...any) (sql.Result, error) {
	return s.db.Exec(query, args...)
}

func (s *SqlHandler) QueryRow(query string, args ...any) *sql.Row {
	return s.db.QueryRow(query, args...)
}

func (s *SqlHandler) Begin() (*sql.Tx, error) {
	return s.db.Begin()
}
