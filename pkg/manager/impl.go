package manager

import (
	"database/sql"

	"code.ysitd.cloud/common/go/db"
)

type twoPhase struct {
	confirm chan bool
	err     chan error
	db      *sql.DB
	tx      *sql.Tx
}
type manager struct {
	app        *appStore
	deployment *deploymentStore
	env        *envStore
	network    *networkStore
}

type appStore struct {
	db db.Pool
}

type deploymentStore struct {
	db db.Pool
}

type envStore struct {
	db db.Pool
}

type networkStore struct {
	db db.Pool
}

type scannable interface {
	Scan(dest ...interface{}) error
}
