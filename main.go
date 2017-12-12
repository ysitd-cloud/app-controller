package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/ysitd-cloud/app-controller/pkg/http"
	"github.com/ysitd-cloud/gin-utils/net"
)

func main() {
	app := gin.Default()
	http.Register(app)
	app.Run(net.GetAddress())
}
