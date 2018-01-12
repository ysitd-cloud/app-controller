package http

import (
	"code.ysitd.cloud/component/deployer/pkg/http/handler"
	"github.com/gin-gonic/gin"
)

func register(app gin.IRouter) {
	group := app.Group("/api/v1")
	registerV1API(group)
}

func registerV1API(app gin.IRoutes) {
	app.POST("/application", handler.CreateApplication)
	app.GET("/user/:user/application", handler.GetApplicationByUsername)
	app.PUT("/application/:app/image", handler.UpdateImage)
	app.GET("/application/:app", handler.GetApplicationById)
}
