package config

import (
	"log"

	"github.com/spf13/viper"

)

var EnvConfigs *envConfigs

func InitiEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

type envConfigs struct {
	LocalServerPort string `mapstructure:"LOCAL_SERVER_PORT"`
	MongoURI        string `mapstructure:"MONGODB_URL"`
	JwtSecret       string `mapstructure:"JWT_SECRET"`
	GEMINI_API_KEY  string  `mapstructure:"GEMINI_API_KEY"`
}

func loadEnvVariables() *envConfigs {
	var config envConfigs

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return &config
}
