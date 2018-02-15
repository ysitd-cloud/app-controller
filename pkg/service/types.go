package service

import (
	"code.ysitd.cloud/component/deployer/pkg/deployer"
	"code.ysitd.cloud/component/deployer/pkg/manager"
)

type service struct {
	manager  manager.Manager
	deployer deployer.Controller
}
