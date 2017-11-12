package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/app-controller/http"
	"github.com/ysitd-cloud/gin-utils/net"
)

func main() {
	app := gin.Default()
	http.Register(app)
	app.Run(net.GetAddress())
}
