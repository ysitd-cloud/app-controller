package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/app-controller/provider"
)

func BindKernel(c *gin.Context) {
	c.Set("kernel", provider.Kernel)
}
