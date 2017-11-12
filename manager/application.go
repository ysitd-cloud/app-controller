package manager

import (
	"database/sql"

	"github.com/satori/go.uuid"
)


func (m *Manager) CreateApplication(app Application) error {
	if app.ID == "" {
		app.ID = uuid.NewV4().String()
	}
	sql := `INSERT INTO application (id, owner) VALUES (?, ?)`
	result, err := m.db.Exec(sql, app.ID, app.Owner)
	if err != nil {
		return err
	}

	if row, err := result.RowsAffected(); err != nil {
		return err
	} else if row != 1 {
		return IncorrectNumOfRowAffected
	}

	if err := m.CreateDeployment(app.ID, app.Deployment); err != nil {
		return err
	}

	if err := m.CreateEnvironment(app.ID, app.Environment); err != nil {
		return err
	}

	return m.CreateNetwork(app.ID, app.Network)
}

func (m *Manager) GetApplicationByID(id string) (*Application, error) {
	query := `SELECT owner FROM application WHERE id = ?`
	row := m.db.QueryRow(query, id)

	var app Application
	var err error
	if err := row.Scan(&app.Owner); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	if app.Deployment, err = m.GetDeployment(id); err != nil {
		return nil, err
	}

	if app.Environment, err = m.GetEnvironment(id); err != nil {
		return nil, err
	}

	if app.Network, err = m.GetNetwork(id); err != nil {
		return nil, err
	}

	app.ID = id

	return &app, nil
}

func (m *Manager) GetApplicationByOwner(owner string) ([]*Application, error) {
	query := `SELECT id FROM application WHERE owner = ?`
	rows, err := m.db.Query(query, owner)
	if err != nil {
		return nil, err
	}

	apps := make([]*Application, 0)

	for rows.Next() {
		var app Application
		var err error
		if err := rows.Scan(&app.ID); err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			return nil, err
		}

		id := app.ID

		if app.Deployment, err = m.GetDeployment(id); err != nil {
			return nil, err
		}

		if app.Environment, err = m.GetEnvironment(id); err != nil {
			return nil, err
		}

		if app.Network, err = m.GetNetwork(id); err != nil {
			return nil, err
		}

		app.Owner = owner

		apps = append(apps, &app)
	}

	return apps, nil
}

func (m *Manager) DeleteApplication(id string) (error) {
	var err error
	if err = m.DeleteNetwork(id); err != nil {
		return err
	}

	if err = m.DeleteEnvironment(id); err != nil {
		return err
	}

	if err = m.DeleteDeployment(id); err != nil {
		return err
	}

	sql := `DELETE FROM application WHERE id = $1`
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
