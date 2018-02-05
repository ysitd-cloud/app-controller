package manager

import (
	"code.ysitd.cloud/common/go/db"
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
