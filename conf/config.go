package conf

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_PORT        string
	DB_HOST        string
	DB_NAME        string
	REDIS_ADDR     string
	REDIS_PASSWORD string
	REDIS_BD       string
	REDIS_PORT     string
}

func GetConfig() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("config/config.json", &configuration)
	return configuration
}
