package manager

import (
	"database/sql"

	"code.ysitd.cloud/grpc/schema/deployer/models"
)

func prepareCreateApplication(app *Application, tx *sql.Tx) (err error) {
	query := `INSERT INTO applications (id, owner, name) VALUES ($1, $2, $3)`
	_, err = tx.Exec(query, app.ID, app.Owner, app.Name)
	return
}

func prepareDeleteApplication(id string, tx *sql.Tx) (err error) {
	query := `DELETE FROM applications WHERE id = $1`
	_, err = tx.Exec(query, id)
	return
}

func prepareCreateDeployment(id string, deployment *models.Deployment, tx *sql.Tx) (err error) {
	query := `INSERT INTO app_deployment (app, image, tag) VALUES ($1, $2, $3)`
	image := deployment.Image
	tag := deployment.Tag
	_, err = tx.Exec(query, id, image, tag)
	return
}

func prepareUpdateDeployment(id string, deployment *models.Deployment, tx *sql.Tx) (err error) {
	query := `UPDATE app_deployment SET image = $2, tag = $3 WHERE app = $1`
	image := deployment.Image
	tag := deployment.Tag
	_, err = tx.Exec(query, id, image, tag)
	return
}

func prepareDeleteDeployment(id string, tx *sql.Tx) (err error) {
	query := `DELETE FROM app_deployment WHERE app = $1`
	_, err = tx.Exec(query, id)
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

func prepareUpdateEnvironment(id string, env Environment, tx *sql.Tx) (err error) {
	values, err := json.Marshal(env)
	if err != nil {
		return
	}
	query := `UPDATE app_environment SET values = $2 WHERE app = $1`
	_, err = tx.Exec(query, id, string(values))
	return
}

func prepareDeleteEnvironment(id string, tx *sql.Tx) (err error) {
	query := `DELETE FROM app_environment WHERE app = $1`
	_, err = tx.Exec(query, id)
	return
}

func prepareCreateNetwork(id string, network *models.Network, tx *sql.Tx) (err error) {
	query := `INSERT INTO app_network (app, domain) VALUES ($1, $2)`
	_, err = tx.Exec(query, id, network.Domain)
	return
}

func prepareUpdateNetwork(id string, network *models.Network, tx *sql.Tx) (err error) {
	query := `UPDATE app_network SET domain = $2 WHERE app = $1`
	_, err = tx.Exec(query, id, network.Domain)
	return
}

func prepareDeleteNetwork(id string, tx *sql.Tx) (err error) {
	query := `DELETE FROM app_network WHERE app = $1`
	_, err = tx.Exec(query, id)
	return
}
