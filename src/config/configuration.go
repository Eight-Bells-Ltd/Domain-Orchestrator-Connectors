package config

import (
	"log"
	"github.com/spf13/viper"

	"doc/src/data"
)

var Config *data.Configuration

func SetupConfiguration() {
	var configuration *data.Configuration

	//ToDo SetUp configuration file hardcode
	viper.SetConfigFile("src/config/config.yml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	Config = configuration
}

// GetConfig provides configuration data
func GetConfig() *data.Configuration {
	return Config
}
