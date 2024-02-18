package configs

import (
	"github.com/spf13/viper"
	"strconv"
)

type ApiConf struct {
	Port int
}

type config struct {
	ApiConf *ApiConf
}

var Configs *config

func InitConfigs(configFile string) {
	InitEnv(configFile)
	Configs = new(config)
	Configs.ApiConf = buildApiConfig()
}

func buildApiConfig() *ApiConf {
	var port, _ = strconv.Atoi(viper.Get("PORT").(string))
	return &ApiConf{Port: port}
}
