package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config *Configuration

func SetupConfiguration() {
	var configuration *Configuration

	//ToDo SetUp configuration file hardcode
	viper.SetConfigFile("config")
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
func GetConfig() *Configuration {
	return Config
}
