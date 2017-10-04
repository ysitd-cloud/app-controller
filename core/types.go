package core

import (
	"github.com/Azure/azure-storage-go"
	"github.com/ysitd-cloud/app-controller/app"
)

type Manager interface {
	GetApplication(id string) app.Application
}

type cEnvironmentManager interface {
	GetEntry(id string) app.Environment
}

type environmentManager struct {
	client storage.Table
}

type MetaInformationManager interface {
	GetEntity(id string) app.MetaInformation
}

type AutoScaleManager interface {
	GetEntry(id string) app.AutoScale
}

type NetworkManager interface {
	GetEntry(id string) app.Network
}
