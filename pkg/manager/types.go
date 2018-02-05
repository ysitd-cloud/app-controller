package manager

import (
	"database/sql"

	"code.ysitd.cloud/common/go/db"
	"code.ysitd.cloud/grpc/schema/deployer/models"
)

type Manager interface {
	SetDB(db db.Pool)

	GetApplicationStore() ApplicationStore
	GetDeploymentStore() DeploymentStore
	GetEnvironmentStore() EnvironmentStore
	GetNetworkStore() NetworkStore

	CreateApplication(app *Application) (confirm chan<- bool, e <-chan error, err error)
	GetApplicationByID(id string) (*Application, error)
	GetApplicationByOwner(owner string) ([]*Application, error)
	DeleteApplication(id string) (confirm chan<- bool, e <-chan error, err error)

	GetDeployment(id string) (*models.Deployment, error)
	CreateDeployment(id string, deployment *models.Deployment) (confirm chan<- bool, e <-chan error, err error)
	UpdateDeployment(id string, deployment *models.Deployment) (confirm chan<- bool, e <-chan error, err error)
	DeleteDeployment(id string) (confirm chan<- bool, e <-chan error, err error)

	GetEnvironment(id string) (Environment, error)
	CreateEnvironment(id string, env Environment) (confirm chan<- bool, e <-chan error, err error)
	UpdateEnvironment(id string, env Environment) (confirm chan<- bool, e <-chan error, err error)
	DeleteEnvironment(id string) (confirm chan<- bool, e <-chan error, err error)

	GetNetwork(id string) (*models.Network, error)
	CreateNetwork(id string, network *models.Network) (confirm chan<- bool, e <-chan error, err error)
	UpdateNetwork(id string, network *models.Network) (confirm chan<- bool, e <-chan error, err error)
	DeleteNetwork(id string) (confirm chan<- bool, e <-chan error, err error)
}

type TwoPhaseConfirm interface {
	Ok() (err error)
	Cancel() (err error)
	GetTx() *sql.Tx
}

type ApplicationStore interface {
	Create(app *Application) (TwoPhaseConfirm, error)
	GetByID(id string) (*Application, error)
	GetByOwner(owner string) ([]*Application, error)
	Delete(id string) (TwoPhaseConfirm, error)
}

type DeploymentStore interface {
	Get(id string) (*models.Deployment, error)
	Create(id string, deployment *models.Deployment) (TwoPhaseConfirm, error)
	Update(id string, deployment *models.Deployment) (TwoPhaseConfirm, error)
	Delete(id string) (TwoPhaseConfirm, error)
}

type EnvironmentStore interface {
	Get(id string) (Environment, error)
	Create(id string, env Environment) (TwoPhaseConfirm, error)
	Update(id string, env Environment) (TwoPhaseConfirm, error)
	Delete(id string) (TwoPhaseConfirm, error)
}

type NetworkStore interface {
	Get(id string) (*models.Network, error)
	Create(id string, network *models.Network) (TwoPhaseConfirm, error)
	Update(id string, network *models.Network) (TwoPhaseConfirm, error)
	Delete(id string) (TwoPhaseConfirm, error)
}

type Application struct {
	ID          string             `json:"id,omitempty"`
	Owner       string             `json:"owner"`
	Name        string             `json:"name"`
	Deployment  *models.Deployment `json:"deployment"`
	Environment Environment        `json:"environment"`
	Network     *models.Network    `json:"network"`
}

type Environment map[string]string
