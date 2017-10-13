package core

import "github.com/ysitd-cloud/app-controller/app"

func NewManager(
	env EnvironmentManager,
	meta MetaInformationManager,
	autoScale AutoScaleManager,
	network NetworkManager,
) Manager {
	return &manager{
		env:       env,
		meta:      meta,
		autoScale: autoScale,
		network:   network,
	}
}

func (m *manager) GetApplication(id string) app.Application {
	env := m.env.GetEntry(id)
	info := m.meta.GetEntity(id)
	autoScale := m.autoScale.GetEntry(id)
	network := m.network.GetEntry(id)

	return app.NewApplication(id, env, info, autoScale, network)
}

func (m *manager) Close() {
	m.env.Close()
	m.meta.Close()
	m.autoScale.Close()
	m.network.Close()
}
