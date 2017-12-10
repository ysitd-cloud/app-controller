package manager

import (
	"database/sql"
	"encoding/json"
)

func (m *Manager) GetEnvironment(id string) (Environment, error) {
	query := `SELECT values FROM app_environment WHERE app = $1`
	row := m.db.QueryRow(query, id)

	var values string

	if err := row.Scan(&values); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	env := make(map[string]string)
	json.Unmarshal([]byte(values), env)

	return env, nil
}

func (m *Manager) CreateEnvironment(id string, env Environment) (error) {
	values, err := json.Marshal(env)
	if err != nil {
		return err
	}
	query := `INSERT INTO app_environment (app, values) VALUES ($1, $2)`
	result, err := m.db.Exec(query, id, string(values))
	if err != nil {
		return err
	}

	if row, err := result.RowsAffected(); err != nil {
		return err
	} else if row != 1 {
		return IncorrectNumOfRowAffected
	}

	return nil
}

func (m *Manager) UpdateEnvironment(id string, env Environment) error {
	values, err := json.Marshal(env)
	if err != nil {
		return err
	}

	sql := `UPDATE app_environment SET values = $2 WHERE app = $1`
	result, err := m.db.Exec(sql, id, values)
	if err != nil {
		return err
	}

	if row, err := result.RowsAffected(); err != nil {
		return err
	} else if row != 1 {
		return IncorrectNumOfRowAffected
	}

	return nil
}

func (m *Manager) DeleteEnvironment(id string) error {
	sql := `DELETE FROM app_environment WHERE app = $1`
	result, err := m.db.Exec(sql, id)
	if err != nil {
		return err
	}

	if row, err := result.RowsAffected(); err != nil {
		return err
	} else if row != 1 {
		return IncorrectNumOfRowAffected
	}

	return nil
}
