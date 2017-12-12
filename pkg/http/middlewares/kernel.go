package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/app-controller/pkg/kernel"
)

func BindKernel(c *gin.Context) {
	c.Set("kernel", kernel.Kernel)
}
