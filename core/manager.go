package core

import (
	"database/sql"
	"github.com/Azure/azure-storage-go"
)

type AppManager struct {
	DB *sql.DB
	Table *storage.Table
}

func NewAppManager(db *sql.DB, table *storage.Table) *AppManager {
	return &AppManager{
		DB: db,
		Table: table,
	}
}
