package manager

import (
	"database/sql"
	"github.com/ysitd-cloud/app-controller/app"
)

type EnvironmentManager struct {
	db *sql.DB
}

func (manager *EnvironmentManager) Get(id string) (app.Environment, error) {
	query := "SELECT key, value FROM app_env WHERE id = $1"
	rows, err := manager.db.Query(query, id)

	env := make(map[string]string)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var key string
		var value string
		if err := rows.Scan(&key, &value); err != nil {
			return nil, err
		}
		env[key] = value

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return env, nil
}
