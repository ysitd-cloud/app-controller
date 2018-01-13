package main

import (
	"net"

	"code.ysitd.cloud/component/deployer/pkg/kernel"
	ginNet "code.ysitd.cloud/gin/utils/net"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

func main() {
	{
		app := kernel.Kernel.Make("http.server").(*gin.Engine)
		go app.Run(ginNet.GetAddress())
	}

	{
		server := kernel.Kernel.Make("grpc.server").(*grpc.Server)
		listener := kernel.Kernel.Make("grpc.listener").(net.Listener)
		server.Serve(listener)
	}
}
