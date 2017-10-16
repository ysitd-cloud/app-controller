package core

import (
	"database/sql"

	"github.com/Azure/azure-storage-go"
	"github.com/ysitd-cloud/app-controller/app"
	appv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	extv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
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

type KubernetesDeployer struct {
	deployment appv1beta1.DeploymentInterface
	ingress    extv1beta1.IngressInterface
	service    corev1.ServiceInterface
	secret     corev1.SecretInterface
	manager    Manager
}
