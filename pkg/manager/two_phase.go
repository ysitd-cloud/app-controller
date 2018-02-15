package manager

import (
	"database/sql"

	"code.ysitd.cloud/common/go/db"
)

func (t *twoPhase) start(pool db.Pool) (err error) {
	db, err := pool.Acquire()
	if err != nil {
		return
	}
	tx, err := db.Begin()
	if err != nil {
		return
	}

	t.db = db
	t.tx = tx
	t.confirm = make(chan bool)
	t.err = make(chan error)
	return
}

func (t *twoPhase) GetTx() *sql.Tx {
	return t.tx
}

func (t *twoPhase) Ok() (err error) {
	err = t.tx.Commit()
	t.cleanup()
	return
}

func (t *twoPhase) Cancel() (err error) {
	err = t.tx.Rollback()
	t.cleanup()
	return
}

func (t *twoPhase) cleanup() {
	defer t.db.Close()
	defer close(t.err)
}
