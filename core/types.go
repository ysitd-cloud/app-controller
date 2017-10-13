package core

import (
	"database/sql"

	"github.com/Azure/azure-storage-go"
	"github.com/ysitd-cloud/app-controller/app"
)

type Manager interface {
	GetApplication(id string) app.Application
	Close()
}

type manager struct {
	env       EnvironmentManager
	meta      MetaInformationManager
	autoScale AutoScaleManager
	network   NetworkManager
}

type EnvironmentManager interface {
	GetEntry(id string) app.Environment
	Close()
}

type environmentManager struct {
	client storage.Table
}

type MetaInformationManager interface {
	GetEntity(id string) app.MetaInformation
	Close()
}

type metaInformationManager struct {
	db *sql.DB
}

type AutoScaleManager interface {
	GetEntry(id string) app.AutoScale
	Close()
}

type noAutoScale struct{}

type NetworkManager interface {
	GetEntry(id string) app.Network
	Close()
}

type networkManager struct {
	db *sql.DB
}
