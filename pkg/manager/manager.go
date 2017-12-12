package manager

import (
	"database/sql"
	"encoding/json"

	"github.com/satori/go.uuid"
)

func (m *manager) SetDB(db *sql.DB) {
	m.db = db
}

func (m *manager) Close() {
	m.db.Close()
}

func (m *manager) CreateApplication(app Application) error {
	if app.ID == "" {
		app.ID = uuid.NewV4().String()
	}
	query := `INSERT INTO applications (id, owner, name) VALUES ($1, $2, $3)`
	result, err := m.db.Exec(query, app.ID, app.Owner, app.Name)
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

func (m *manager) GetApplicationByID(id string) (*Application, error) {
	query := `SELECT owner, name FROM applications WHERE id = $1`
	row := m.db.QueryRow(query, id)

	var app Application
	var err error
	if err := row.Scan(&app.Owner, &app.Name); err == sql.ErrNoRows {
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

func (m *manager) GetApplicationByOwner(owner string) ([]*Application, error) {
	query := `SELECT id, name FROM applications WHERE owner = $1`
	rows, err := m.db.Query(query, owner)
	if err != nil {
		return nil, err
	}

	apps := make([]*Application, 0)

	for rows.Next() {
		var app Application
		var err error
		if err := rows.Scan(&app.ID, &app.Name); err == sql.ErrNoRows {
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

func (m *manager) DeleteApplication(id string) error {
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

	sql := `DELETE FROM applications WHERE id = $1`
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

func (m *manager) GetDeployment(id string) (*Deployment, error) {
	query := `SELECT image, tag FROM app_deployment WHERE app = $1`
	row := m.db.QueryRow(query, id)

	var deployment Deployment

	if e := row.Scan(&deployment.Image, &deployment.Tag); e == sql.ErrNoRows {
		return nil, nil
	} else if e != nil {
		return nil, e
	}

	return &deployment, nil
}

func (m *manager) CreateDeployment(id string, deployment *Deployment) error {
	query := `INSERT INTO app_deployment (app, image, tag) VALUES ($1, $2, $3)`
	image := deployment.Image
	tag := deployment.Tag
	result, err := m.db.Exec(query, id, image, tag)
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

func (m *manager) UpdateDeployment(id string, deployment *Deployment) error {
	sql := `UPDATE app_deployment SET image = $2, tag = $3 WHERE app = $1`
	image := deployment.Image
	tag := deployment.Tag
	result, err := m.db.Exec(sql, id, image, tag)
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

func (m *manager) DeleteDeployment(id string) error {
	query := `DELETE FROM app_deployment WHERE app = $1`
	result, err := m.db.Exec(query, id)
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

func (m *manager) GetEnvironment(id string) (Environment, error) {
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

func (m *manager) CreateEnvironment(id string, env Environment) error {
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

func (m *manager) UpdateEnvironment(id string, env Environment) error {
	values, err := json.Marshal(env)
	if err != nil {
		return err
	}

	query := `UPDATE app_environment SET values = $2 WHERE app = $1`
	result, err := m.db.Exec(query, id, values)
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

func (m *manager) DeleteEnvironment(id string) error {
	query := `DELETE FROM app_environment WHERE app = $1`
	result, err := m.db.Exec(query, id)
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

func (m *manager) GetNetwork(id string) (*Network, error) {
	query := `SELECT domain FROM app_network WHERE app = $1`
	row := m.db.QueryRow(query, id)

	var network Network

	if e := row.Scan(&network.Domain); e == sql.ErrNoRows {
		return nil, nil
	} else if e != nil {
		return nil, e
	}

	return &network, nil
}

func (m *manager) CreateNetwork(id string, network *Network) error {
	query := `INSERT INTO app_network (app, domain) VALUES ($1, $2)`
	result, err := m.db.Exec(query, id, network.Domain)
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

func (m *manager) UpdateNetwork(id string, network *Network) error {
	query := `UPDATE app_network SET domain = $2 WHERE app = $1`
	result, err := m.db.Exec(query, id, network.Domain)
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

func (m *manager) DeleteNetwork(id string) error {
	query := `DELETE FROM app_network WHERE app = $1`
	result, err := m.db.Exec(query, id)
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
