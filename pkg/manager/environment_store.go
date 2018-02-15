package manager

import (
	"database/sql"
	"encoding/json"
)

func (es *envStore) Get(id string) (Environment, error) {
	db, err := es.db.Acquire()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `SELECT values FROM app_environment WHERE app = $1`
	row := db.QueryRow(query, id)

	var values string

	if err := row.Scan(&values); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	env := make(map[string]string)
	json.Unmarshal([]byte(values), &env)

	return env, nil
}

func (es *envStore) Create(id string, env Environment) (TwoPhaseConfirm, error) {
	t := new(twoPhase)
	if err := t.start(es.db); err != nil {
		return nil, err
	}

	prepareCreateEnvironment(id, env, t.GetTx())
	return t, nil
}

func (es *envStore) Update(id string, env Environment) (TwoPhaseConfirm, error) {
	t := new(twoPhase)
	if err := t.start(es.db); err != nil {
		return nil, err
	}

	prepareUpdateEnvironment(id, env, t.GetTx())
	return t, nil
}

func (es *envStore) Delete(id string) (TwoPhaseConfirm, error) {
	t := new(twoPhase)
	if err := t.start(es.db); err != nil {
		return nil, err
	}

	prepareDeleteEnvironment(id, t.GetTx())
	return t, nil
}
