package manager

import (
	"database/sql"
)

func (m *Manager) GetDeployment(id string) (*Deployment, error) {
	query := `SELECT image, tag WHERE id = $1`
	row := m.db.QueryRow(query, id)

	var deployment Deployment

	if e := row.Scan(&deployment.Image, &deployment.Tag); e == sql.ErrNoRows {
		return nil, nil
	} else if e != nil {
		return nil, e
	}

	return &deployment, nil
}

func (m *Manager) CreateDeployment(id string, deployment *Deployment) error {
	sql := `INSERT INTO app_deployment (id, image, tag) VALUES (?, ?, ?)`
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

func (m *Manager) UpdateDeployment(id string, deployment *Deployment) error {
	sql := `UPDATE app_deployment SET image = $2, tag = $3 WHERE id = $1`
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

func (m *Manager) DeleteDeployment(id string) error {
	sql := `DELETE FROM app_deployment WHERE id = $1`
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
