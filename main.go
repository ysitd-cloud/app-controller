package main

import (
	"net"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/ysitd-cloud/app-controller/pkg/kernel"
	ginNet "github.com/ysitd-cloud/gin-utils/net"
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
