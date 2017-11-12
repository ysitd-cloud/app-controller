package manager

import (
	"database/sql"
	"errors"
)

var (
	IncorrectNumOfRowAffected = errors.New("incorrect number of row affected")
)

type Manager struct {
	db *sql.DB
}

type Application struct {
	ID          string      `json:"id,omitempty"`
	Owner       string      `json:"owner"`
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
