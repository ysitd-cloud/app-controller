package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/ysitd-cloud/app-controller/manager/register"
	"github.com/ysitd-cloud/gin-utils/net"
)

func main() {
	app := register.Kernel.Make("http.server").(gin.Engine)
	app.Run(net.GetAddress())
}
