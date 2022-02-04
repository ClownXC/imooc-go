package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"my_shop/user-web/global"
)

func GetEnv(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {

	debug := GetEnv("IMOOC_SHOP")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user-web/%s-pro.yml", configFilePrefix)

	// TODO:
	debug = true

	if debug {
		configFileName = fmt.Sprintf("user-web/%s-debug.yml", configFilePrefix)
		zap.S().Infof("====================开发模式====================")
	}

	fmt.Println("configFileName: " + configFileName)
	v := viper.New()
	v.SetConfigFile(configFileName)

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = v.Unmarshal(&global.ServerConfig)
	if err != nil {
		panic(err)
	}

	zap.S().Infof("配置信息： %v", global.ServerConfig)

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置文件发生变化： %s", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&global.ServerConfig)
		zap.S().Infof("配置信息： %v", global.ServerConfig)

	})
}
