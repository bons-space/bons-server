package initialize

import (
	"bons-server/src/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default();
	publicGroup := Router.Group("")
	{
		router.InitUser(publicGroup)
	}
	return Router
}