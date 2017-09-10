package connect

import (
	_ "github.com/lib/pq"
	"database/sql"
)

func NewPostgresClient(connection string) (*sql.DB, error) {
	return sql.Open("postgres", connection)
}

