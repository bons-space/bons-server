package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUser(Router *gin.RouterGroup) (R gin.IRoutes) {
	UserRoute := Router.Group("/")
	{
		UserRoute.GET("login", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": 201,
				"msg": "登录成功",
			})
		})
	}
	return UserRoute
}