package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/app-controller/pkg/http/handler"
	"github.com/ysitd-cloud/app-controller/pkg/http/middlewares"
)

func Register(app *gin.Engine) {
	app.Use(middlewares.BindKernel)
	group := app.Group("/api/v1")
	registerV1API(group)
}

func registerV1API(app *gin.RouterGroup) {
	app.POST("/application", handler.CreateApplication)
	app.GET("/user/:user/application", handler.GetApplicationByUsername)
	app.PUT("/application/:app/image", handler.UpdateImage)
	app.GET("/application/:app", handler.GetApplicationById)
}
