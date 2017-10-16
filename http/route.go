package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/app-controller/http/handler"
)

func Register(app *gin.Engine) {
	group := app.Group("/api/v1")
	registerV1API(group)
}

func registerV1API(app *gin.RouterGroup) {
	app.GET("/app/:id/meta", handler.GetMeta)
	app.GET("/app/:id/network", handler.GetNetwork)
	app.GET("/app/:id/env", handler.GetEnv)
}
