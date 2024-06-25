package config

import (
	"log"
	"github.com/spf13/viper"

	"doc/src/data"
)

var Config *data.AppConfiguration

func SetupConfiguration() {
	var appConfiguration *data.AppConfiguration

	//ToDo SetUp configuration file hardcode
	viper.SetConfigFile("src/config/config.yml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&appConfiguration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	Config = appConfiguration
	log.Printf("App Configuration readed")
}

// GetConfig provides configuration data
func GetConfig() data.AppConfiguration {
	return *Config
}

// GetURL provides URL configured in config.yml
func GetURL() string {
	log.Printf("ENTER: GetURL(), %s", Config.Orchestrator)
	return Config.Orchestrator
}