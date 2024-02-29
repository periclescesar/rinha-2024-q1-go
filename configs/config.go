package configs

import (
	"github.com/spf13/viper"
	"strconv"
)

type ApiConf struct {
	Port int
}

type DbConf struct {
	Host   string
	Port   int
	User   string
	Pass   string
	Dbname string
}

type config struct {
	ApiConf *ApiConf
	DbConf  *DbConf
}

var Configs *config

func InitConfigs(configFile string) {
	InitEnv(configFile)
	Configs = new(config)
	Configs.ApiConf = buildApiConfig()
	Configs.DbConf = buildDbConfig()
}

func buildApiConfig() *ApiConf {
	var port, _ = strconv.Atoi(viper.Get("PORT").(string))
	return &ApiConf{Port: port}
}

func buildDbConfig() *DbConf {
	var port, _ = strconv.Atoi(viper.Get("DB_PORT").(string))
	var host, _ = viper.Get("DB_HOST").(string)
	var user, _ = viper.Get("DB_USER").(string)
	var pass, _ = viper.Get("DB_PASS").(string)
	var dbname, _ = viper.Get("DB_DBNAME").(string)

	return &DbConf{Host: host, Port: port, User: user, Pass: pass, Dbname: dbname}
}
