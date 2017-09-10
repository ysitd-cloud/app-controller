package connect

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func GetDB() (*sql.DB, error) {
	dbUrl := os.Getenv("DB_URL")
	return sql.Open("postgres", dbUrl)
}
