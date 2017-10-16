package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/ysitd-cloud/app-controller/http"
	"github.com/ysitd-cloud/app-controller/kernel"
	"github.com/ysitd-cloud/gin-utils/net"
)

func main() {
	app := kernel.Kernel.Make("http.server").(*gin.Engine)
	http.Register(app)
	app.Run(net.GetAddress())
}
