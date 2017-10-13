package core

import (
	"database/sql"

	"github.com/ysitd-cloud/app-controller/app"
)

func NewMetaInformationManager(db *sql.DB) MetaInformationManager {
	return &metaInformationManager{
		db: db,
	}
}

func (m *metaInformationManager) GetEntity(id string) app.MetaInformation {
	query := "SELECT image, tag FROM app_deploy WhERE app = $1 LIMIT 1"

	row := m.db.QueryRow(query, id)

	var image, tag string
	row.Scan(&image, &tag)

	return app.NewMetaInformation(image, tag)
}

func (m *metaInformationManager) Close() {
	m.db.Close()
}
