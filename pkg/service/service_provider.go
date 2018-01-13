package service

import (
	"net"

	"code.ysitd.cloud/component/deployer/pkg/deployer"
	"code.ysitd.cloud/component/deployer/pkg/manager"
	pb "code.ysitd.cloud/grpc/schema/deployer"
	"github.com/tonyhhyip/go-di-container"
	"google.golang.org/grpc"
)

type grpcServiceProvider struct {
	*container.AbstractServiceProvider
}

func (*grpcServiceProvider) Provides() []string {
	return []string{
		"service",
		"grpc.server",
	}
}

func (*grpcServiceProvider) Register(app container.Container) {
	app.Bind("service", func(app container.Container) interface{} {
		m := app.Make("manager").(manager.Manager)
		d := app.Make("deployer").(deployer.Controller)
		return newService(m, d)
	})

	app.Singleton("grpc.server", func(app container.Container) interface{} {
		server := grpc.NewServer()
		service := app.Make("service").(pb.DeployerServer)
		pb.RegisterDeployerServer(server, service)
		return server
	})

	app.Singleton("grpc.listener", func(app container.Container) interface{} {
		listener, err := net.Listen("tcp", "localhost:50051")
		if err != nil {
			panic(err)
		}

		return listener
	})
}
