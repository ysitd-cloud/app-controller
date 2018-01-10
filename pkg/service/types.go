package service

import (
	"github.com/ysitd-cloud/app-controller/pkg/deployer"
	"github.com/ysitd-cloud/app-controller/pkg/manager"
)

type service struct {
	manager  manager.Manager
	deployer deployer.Controller
}
