package manager

import (
	"database/sql"
)

func (m *Manager) GetNetwork(id string) (*Network, error) {
	query := `SELECT domain WHERE id = $1`
	row := m.db.QueryRow(query, id)

	var network Network

	if e := row.Scan(&network.Domain); e == sql.ErrNoRows {
		return nil, nil
	} else if e != nil {
		return nil, e
	}

	return &network, nil
}

func (m *Manager) CreateNetwork(id string, network *Network) error {
	sql := `INSERT INTO app_network (id, domain) VALUES (?, ?)`
	result, err := m.db.Exec(sql, id, network.Domain)
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

func (m *Manager) UpdateNetwork(id string, network *Network) error {
	sql := `UPDATE app_network SET domain = $2 WHERE id = $1`
	result, err := m.db.Exec(sql, id, network.Domain)
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

func (m *Manager) DeleteNetwork(id string) error {
	sql := `DELETE FROM app_network WHERE id = $1`
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
