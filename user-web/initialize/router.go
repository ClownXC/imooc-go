package initialize

import (
	"github.com/gin-gonic/gin"
	"my_shop/user-web/middlewares"
	router2 "my_shop/user-web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/v1")
	router2.InitUserRouter(ApiGroup)

	return Router
}
