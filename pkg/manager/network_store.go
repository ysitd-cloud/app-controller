package manager

import (
	"database/sql"

	"code.ysitd.cloud/grpc/schema/deployer/models"
)

func (ns *networkStore) Get(id string) (*models.Network, error) {
	db, err := ns.db.Acquire()
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

func (ns *networkStore) Create(id string, network *models.Network) (TwoPhaseConfirm, error) {
	t := new(twoPhase)
	if err := t.start(ns.db); err != nil {
		return nil, err
	}
	prepareUpdateNetwork(id, network, t.GetTx())
	return t, nil
}

func (ns *networkStore) Update(id string, network *models.Network) (TwoPhaseConfirm, error) {
	t := new(twoPhase)
	if err := t.start(ns.db); err != nil {
		return nil, err
	}
	prepareUpdateNetwork(id, network, t.GetTx())
	return t, nil
}

func (ns *networkStore) Delete(id string) (TwoPhaseConfirm, error) {
	t := new(twoPhase)
	if err := t.start(ns.db); err != nil {
		return nil, err
	}

	prepareDeleteNetwork(id, t.GetTx())
	return t, nil
}
