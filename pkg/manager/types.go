package manager

import (
	"errors"

	"github.com/ysitd-cloud/go-common/db"
)

var (
	IncorrectNumOfRowAffected = errors.New("incorrect number of row affected")
)

type Manager interface {
	SetDB(db db.Pool)

	CreateApplication(app Application) (confirm chan<- bool, e <-chan error, err error)
	GetApplicationByID(id string) (*Application, error)
	GetApplicationByOwner(owner string) ([]*Application, error)
	DeleteApplication(id string) (confirm chan<- bool, e <-chan error, err error)

	GetDeployment(id string) (*Deployment, error)
	CreateDeployment(id string, deployment *Deployment) (confirm chan<- bool, e <-chan error, err error)
	UpdateDeployment(id string, deployment *Deployment) (confirm chan<- bool, e <-chan error, err error)
	DeleteDeployment(id string) (confirm chan<- bool, e <-chan error, err error)

	GetEnvironment(id string) (Environment, error)
	CreateEnvironment(id string, env Environment) (confirm chan<- bool, e <-chan error, err error)
	UpdateEnvironment(id string, env Environment) (confirm chan<- bool, e <-chan error, err error)
	DeleteEnvironment(id string) (confirm chan<- bool, e <-chan error, err error)

	GetNetwork(id string) (*Network, error)
	CreateNetwork(id string, network *Network) (confirm chan<- bool, e <-chan error, err error)
	UpdateNetwork(id string, network *Network) (confirm chan<- bool, e <-chan error, err error)
	DeleteNetwork(id string) (confirm chan<- bool, e <-chan error, err error)
}

type manager struct {
	db db.Pool
}

type Application struct {
	ID          string      `json:"id,omitempty"`
	Owner       string      `json:"owner"`
	Name        string      `json:"name"`
	Deployment  *Deployment `json:"deployment"`
	Environment Environment `json:"environment"`
	Network     *Network    `json:"network"`
}

type Deployment struct {
	Image string `json:"image"`
	Tag   string `json:"tag"`
}

type Environment map[string]string

type Network struct {
	Domain string `json:"domain"`
}
