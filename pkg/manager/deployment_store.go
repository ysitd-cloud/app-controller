package manager

import (
	"database/sql"

	"code.ysitd.cloud/grpc/schema/deployer/models"
)

func (ds *deploymentStore) Get(id string) (*models.Deployment, error) {
	db, err := ds.db.Acquire()
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

func (ds *deploymentStore) Create(id string, deployment *models.Deployment) (TwoPhaseConfirm, error) {
	t := new(twoPhase)
	if err := t.start(ds.db); err != nil {
		return nil, err
	}

	prepareCreateDeployment(id, deployment, t.GetTx())
	return t, nil
}

func (ds *deploymentStore) Update(id string, deployment *models.Deployment) (TwoPhaseConfirm, error) {
	t := new(twoPhase)
	if err := t.start(ds.db); err != nil {
		return nil, err
	}

	prepareUpdateDeployment(id, deployment, t.GetTx())
	return t, nil
}

func (ds *deploymentStore) Delete(id string) (TwoPhaseConfirm, error) {
	t := new(twoPhase)
	if err := t.start(ds.db); err != nil {
		return nil, err
	}

	prepareDeleteDeployment(id, t.GetTx())
	return t, nil
}
