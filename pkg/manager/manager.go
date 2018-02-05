package manager

import (
	"database/sql"
	"encoding/json"

	"code.ysitd.cloud/common/go/db"
	"code.ysitd.cloud/grpc/schema/deployer/models"
)

func (m *manager) SetDB(db db.Pool) {
	m.db = db

	if m.app == nil {
		m.app = new(appStore)
	}
	m.app.db = db

	if m.deployment == nil {
		m.deployment = new(deploymentStore)
	}
	m.deployment.db = db

	if m.env == nil {
		m.env = new(envStore)
	}
	m.env.db = db

	if m.network == nil {
		m.network = new(networkStore)
	}
	m.network.db = db
}

func (m *manager) GetApplicationStore() ApplicationStore {
	return m.app
}

func (m *manager) GetDeploymentStore() DeploymentStore {
	return m.deployment
}

func (m *manager) GetEnvironmentStore() EnvironmentStore {
	return m.env
}

func (m *manager) GetNetworkStore() NetworkStore {
	return m.network
}

func (m *manager) init() (db *sql.DB, tx *sql.Tx, confirm chan bool, e chan error, err error) {
	db, err = m.db.Acquire()
	if err != nil {
		return
	}
	tx, err = db.Begin()
	if err != nil {
		return
	}
	confirm = make(chan bool)
	e = make(chan error)
	return
}

func (m *manager) final(db *sql.DB, tx *sql.Tx, confirm <-chan bool, e chan<- error) {
	defer db.Close()
	confirmed := <-confirm

	if confirmed {
		e <- tx.Commit()
	} else {
		e <- tx.Rollback()
	}
	defer close(e)
}

func (m *manager) CreateApplication(app *Application) (confirm chan<- bool, e <-chan error, err error) {
	app = normalizeApplication(app)

	conn, tx, confirmChannel, eChannel, err := m.init()

	prepareCreateApplication(app, tx)
	prepareCreateDeployment(app.ID, app.Deployment, tx)
	prepareCreateEnvironment(app.ID, app.Environment, tx)
	prepareCreateNetwork(app.ID, app.Network, tx)

	go m.final(conn, tx, confirmChannel, eChannel)
	confirm = confirmChannel
	e = eChannel
	return
}

func prepareCreateApplication(app *Application, tx *sql.Tx) (err error) {
	query := `INSERT INTO applications (id, owner, name) VALUES ($1, $2, $3)`
	_, err = tx.Exec(query, app.ID, app.Owner, app.Name)
	return
}

func (m *manager) GetApplicationByID(id string) (*Application, error) {
	db, err := m.db.Acquire()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `SELECT owner, name FROM applications WHERE id = $1`
	row := db.QueryRow(query, id)

	var app Application
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
	db, err := m.db.Acquire()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `SELECT id, name FROM applications WHERE owner = $1`
	rows, err := db.Query(query, owner)
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

func (m *manager) DeleteApplication(id string) (confirm chan<- bool, e <-chan error, err error) {
	conn, tx, confirmChannel, eChannel, err := m.init()

	prepareDeleteNetwork(id, tx)
	prepareDeleteEnvironment(id, tx)
	prepareDeleteDeployment(id, tx)
	prepareDeleteApplication(id, tx)

	go m.final(conn, tx, confirmChannel, eChannel)
	confirm = confirmChannel
	e = eChannel
	return
}

func prepareDeleteApplication(id string, tx *sql.Tx) (err error) {
	query := `DELETE FROM applications WHERE id = $1`
	_, err = tx.Exec(query, id)
	return
}

func (m *manager) GetDeployment(id string) (*models.Deployment, error) {
	db, err := m.db.Acquire()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `SELECT image, tag FROM app_deployment WHERE app = $1`
	row := db.QueryRow(query, id)

	var deployment models.Deployment

	if e := row.Scan(&deployment.Image, &deployment.Tag); e == sql.ErrNoRows {
		return nil, nil
	} else if e != nil {
		return nil, e
	}

	return &deployment, nil
}

func (m *manager) CreateDeployment(id string, deployment *models.Deployment) (confirm chan<- bool, e <-chan error, err error) {
	conn, tx, confirmChannel, eChannel, err := m.init()
	prepareCreateDeployment(id, deployment, tx)
	go m.final(conn, tx, confirmChannel, eChannel)
	confirm = confirmChannel
	e = eChannel
	return
}

func prepareCreateDeployment(id string, deployment *models.Deployment, tx *sql.Tx) (err error) {
	query := `INSERT INTO app_deployment (app, image, tag) VALUES ($1, $2, $3)`
	image := deployment.Image
	tag := deployment.Tag
	_, err = tx.Exec(query, id, image, tag)
	return
}

func (m *manager) UpdateDeployment(id string, deployment *models.Deployment) (confirm chan<- bool, e <-chan error, err error) {
	conn, tx, confirmChannel, eChannel, err := m.init()
	prepareUpdateDeployment(id, deployment, tx)
	go m.final(conn, tx, confirmChannel, eChannel)
	confirm = confirmChannel
	e = eChannel
	return
}

func prepareUpdateDeployment(id string, deployment *models.Deployment, tx *sql.Tx) (err error) {
	query := `UPDATE app_deployment SET image = $2, tag = $3 WHERE app = $1`
	image := deployment.Image
	tag := deployment.Tag
	_, err = tx.Exec(query, id, image, tag)
	return
}

func (m *manager) DeleteDeployment(id string) (confirm chan<- bool, e <-chan error, err error) {
	conn, tx, confirmChannel, eChannel, err := m.init()
	prepareDeleteDeployment(id, tx)
	go m.final(conn, tx, confirmChannel, eChannel)
	confirm = confirmChannel
	e = eChannel
	return
}

func prepareDeleteDeployment(id string, tx *sql.Tx) (err error) {
	query := `DELETE FROM app_deployment WHERE app = $1`
	_, err = tx.Exec(query, id)
	return
}

func (m *manager) GetEnvironment(id string) (Environment, error) {
	db, err := m.db.Acquire()
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

func (m *manager) CreateEnvironment(id string, env Environment) (confirm chan<- bool, e <-chan error, err error) {
	conn, tx, confirmChannel, eChannel, err := m.init()
	prepareCreateEnvironment(id, env, tx)
	go m.final(conn, tx, confirmChannel, eChannel)
	confirm = confirmChannel
	e = eChannel
	return
}

func prepareCreateEnvironment(id string, env Environment, tx *sql.Tx) (err error) {
	values, err := json.Marshal(env)
	if err != nil {
		return
	}
	query := `INSERT INTO app_environment (app, values) VALUES ($1, $2)`
	_, err = tx.Exec(query, id, string(values))
	return
}

func (m *manager) UpdateEnvironment(id string, env Environment) (confirm chan<- bool, e <-chan error, err error) {
	conn, tx, confirmChannel, eChannel, err := m.init()
	prepareUpdateEnvironment(id, env, tx)
	go m.final(conn, tx, confirmChannel, eChannel)
	confirm = confirmChannel
	e = eChannel
	return
}

func prepareUpdateEnvironment(id string, env Environment, tx *sql.Tx) (err error) {
	values, err := json.Marshal(env)
	if err != nil {
		return
	}
	query := `UPDATE app_environment SET values = $2 WHERE app = $1`
	_, err = tx.Exec(query, id, string(values))
	return
}

func (m *manager) DeleteEnvironment(id string) (confirm chan<- bool, e <-chan error, err error) {
	conn, tx, confirmChannel, eChannel, err := m.init()
	prepareDeleteEnvironment(id, tx)
	go m.final(conn, tx, confirmChannel, eChannel)
	confirm = confirmChannel
	e = eChannel
	return
}

func prepareDeleteEnvironment(id string, tx *sql.Tx) (err error) {
	query := `DELETE FROM app_environment WHERE app = $1`
	_, err = tx.Exec(query, id)
	return
}

func (m *manager) GetNetwork(id string) (*models.Network, error) {
	db, err := m.db.Acquire()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `SELECT domain FROM app_network WHERE app = $1`
	row := db.QueryRow(query, id)

	var network models.Network

	if e := row.Scan(&network.Domain); e == sql.ErrNoRows {
		return nil, nil
	} else if e != nil {
		return nil, e
	}

	return &network, nil
}

func (m *manager) CreateNetwork(id string, network *models.Network) (confirm chan<- bool, e <-chan error, err error) {
	conn, tx, confirmChannel, eChannel, err := m.init()
	prepareUpdateNetwork(id, network, tx)
	go m.final(conn, tx, confirmChannel, eChannel)
	confirm = confirmChannel
	e = eChannel
	return
}

func prepareCreateNetwork(id string, network *models.Network, tx *sql.Tx) (err error) {
	query := `INSERT INTO app_network (app, domain) VALUES ($1, $2)`
	_, err = tx.Exec(query, id, network.Domain)
	return
}

func (m *manager) UpdateNetwork(id string, network *models.Network) (confirm chan<- bool, e <-chan error, err error) {
	conn, tx, confirmChannel, eChannel, err := m.init()
	prepareUpdateNetwork(id, network, tx)
	go m.final(conn, tx, confirmChannel, eChannel)
	confirm = confirmChannel
	e = eChannel
	return
}

func prepareUpdateNetwork(id string, network *models.Network, tx *sql.Tx) (err error) {
	query := `UPDATE app_network SET domain = $2 WHERE app = $1`
	_, err = tx.Exec(query, id, network.Domain)
	return
}

func (m *manager) DeleteNetwork(id string) (confirm chan<- bool, e <-chan error, err error) {
	conn, tx, confirmChannel, eChannel, err := m.init()
	prepareDeleteNetwork(id, tx)
	go m.final(conn, tx, confirmChannel, eChannel)
	confirm = confirmChannel
	e = eChannel
	return
}

func prepareDeleteNetwork(id string, tx *sql.Tx) (err error) {
	query := `DELETE FROM app_network WHERE app = $1`
	_, err = tx.Exec(query, id)
	return
}
