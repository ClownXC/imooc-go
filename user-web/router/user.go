package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"my_shop/user-web/api"
	"my_shop/user-web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	zap.S().Info("配置用户相关的 URL")
	{
		UserRouter.GET("/list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.GET("/pwd_login", api.PassWordLogin)
	}

}
