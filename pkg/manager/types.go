package manager

import (
	"errors"

	"github.com/ysitd-cloud/go-common/db"
	"github.com/ysitd-cloud/grpc-schema/deployer/models"
)

var (
	IncorrectNumOfRowAffected = errors.New("incorrect number of row affected")
)

type Manager interface {
	SetDB(db db.Pool)

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

type manager struct {
	db db.Pool
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
