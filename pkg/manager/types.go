package manager

import (
	"database/sql"
	"errors"
)

var (
	IncorrectNumOfRowAffected = errors.New("incorrect number of row affected")
)

type Manager interface {
	SetDB(db *sql.DB)
	Close()

	CreateApplication(app Application) error
	GetApplicationByID(id string) (*Application, error)
	GetApplicationByOwner(owner string) ([]*Application, error)
	DeleteApplication(id string) error

	GetDeployment(id string) (*Deployment, error)
	CreateDeployment(id string, deployment *Deployment) error
	UpdateDeployment(id string, deployment *Deployment) error
	DeleteDeployment(id string) error

	GetEnvironment(id string) (Environment, error)
	CreateEnvironment(id string, env Environment) error
	UpdateEnvironment(id string, env Environment) error
	DeleteEnvironment(id string) error

	GetNetwork(id string) (*Network, error)
	CreateNetwork(id string, network *Network) error
	UpdateNetwork(id string, network *Network) error
	DeleteNetwork(id string) error
}

type manager struct {
	db *sql.DB
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
