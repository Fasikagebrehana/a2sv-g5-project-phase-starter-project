package config

import "github.com/spf13/viper"

type Database struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Uri      string `mapstructure:"uri"`
	Name     string `mapstructure:"name"`
}
type Email struct {
	EmailKey string `mapstructure:"key"`
}
type Config struct {
	Database Database `mapstructure:"database"`
	Email    Email    `mapstructure:"email"`
	Port     string   `mapstructure:"port"`
	JWTKey   string   `mapstructure:"jwt"`
}

func LoadConfig() (*Config, error) {
	//incase to access config file from the root directory
	viper.AddConfigPath("../")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return &Config{}, err
	}
	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		return &config, err
	}
	return &config, nil
}
