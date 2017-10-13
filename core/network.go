package core

import (
	"database/sql"

	"github.com/ysitd-cloud/app-controller/app"
)

func NewNetworkManager(db *sql.DB) NetworkManager {
	return &networkManager{
		db: db,
	}
}

func (m *networkManager) GetEntry(id string) app.Network {
	query := "SELECT port, hostname FROM app_service WHERE app = $1 LIMIT 1"
	row := m.db.QueryRow(query, id)

	var port int32
	var domain string
	row.Scan(&port, &domain)

	return app.NewNetwork(domain, port)
}

func (m *networkManager) Close() {
	m.db.Close()
}
