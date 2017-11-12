package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/app-controller/http/handler"
	"github.com/ysitd-cloud/app-controller/http/middlewares"
)

func Register(app *gin.Engine) {
	app.Use(middlewares.BindKernel)
	group := app.Group("/api/v1")
	registerV1API(group)
}

func registerV1API(app *gin.RouterGroup) {
	app.POST("/application", handler.CreateApplication)
}
