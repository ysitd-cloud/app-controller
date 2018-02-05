package manager

import (
	"database/sql"
	"encoding/json"

	"code.ysitd.cloud/grpc/schema/deployer/models"
)

func (as *appStore) Create(app *Application) (TwoPhaseConfirm, error) {
	app = normalizeApplication(app)

	t := new(twoPhase)
	if err := t.start(as.db); err != nil {
		return nil, err
	}

	prepareCreateApplication(app, t.GetTx())
	prepareCreateDeployment(app.ID, app.Deployment, t.GetTx())
	prepareCreateEnvironment(app.ID, app.Environment, t.GetTx())
	prepareCreateNetwork(app.ID, app.Network, t.GetTx())
	return t, nil
}

func (as *appStore) GetByID(id string) (*Application, error) {
	db, err := as.db.Acquire()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
		SELECT
			applications.id, applications.owner, applications.name,
			app_deployment.image, app_deployment.tag,
			app_environment.values,
			app_network.domain
		FROM applications
		INNER JOIN app_deployment ON app_deployment.app = applications.id
		INNER JOIN app_environment ON app_environment.app = applications.id
		INNER JOIN app_network ON app_network.app = applications.id
		WHERE id = $1
	`
	row := db.QueryRow(query, id)

	return as.scanToApplication(row)
}

func (as *appStore) GetByOwner(owner string) ([]*Application, error) {
	db, err := as.db.Acquire()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
		SELECT
			applications.id, applications.owner, applications.name,
			app_deployment.image, app_deployment.tag,
			app_environment.values,
			app_network.domain
		FROM applications
		INNER JOIN app_deployment ON app_deployment.app = applications.id
		INNER JOIN app_environment ON app_environment.app = applications.id
		INNER JOIN app_network ON app_network.app = applications.id
		WHERE owner = $1
	`
	rows, err := db.Query(query, owner)
	if err != nil {
		return nil, err
	}

	apps := make([]*Application, 0)

	for rows.Next() {
		if app, err := as.scanToApplication(rows); err != nil {
			return nil, err
		} else {
			apps = append(apps, app)
		}
	}

	return apps, nil
}

func (as *appStore) scanToApplication(scanner scannable) (*Application, error) {
	var app Application
	var deployment models.Deployment
	var envString string
	var network models.Network
	if err := scanner.Scan(
		&app.ID,
		&app.Name,
		&app.Owner,
		&app.Name,
		&deployment.Image,
		&deployment.Tag,
		&envString,
		&network.Domain,
	); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	env := make(map[string]string)
	json.Unmarshal([]byte(envString), &env)

	app.Deployment = &deployment
	app.Environment = env
	app.Network = &network

	return &app, nil
}

func (as *appStore) Delete(id string) (TwoPhaseConfirm, error) {
	t := new(twoPhase)
	if err := t.start(as.db); err != nil {
		return nil, err
	}

	prepareDeleteNetwork(id, t.GetTx())
	prepareDeleteEnvironment(id, t.GetTx())
	prepareDeleteDeployment(id, t.GetTx())
	prepareDeleteApplication(id, t.GetTx())
	return t, nil
}
