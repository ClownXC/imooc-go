package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"my_shop/user-web/global"
	"my_shop/user-web/initialize"
	myvaildator "my_shop/user-web/validator"
)

func main() {

	initialize.InitLogger()
	initialize.InitConfig()

	routers := initialize.Routers()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvaildator.ValidateMobile)
	}

	zap.S().Infof("启动服务器端口: %d", global.ServerConfig.Port)
	routers.Run(fmt.Sprintf(":%d", global.ServerConfig.Port))
}
