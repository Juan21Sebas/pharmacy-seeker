package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Server struct {
	Port string `json:"port"`
}

type Api struct {
	Pharmacy Pharmacy
}

type Pharmacy struct {
	Url string
}

type Config struct {
	Server Server
	Api    Api
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigName("properties")
	viper.AddConfigPath("shared/config")
	viper.SetConfigType("yml")

	err = viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	// Set undefined variables
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("/api/v1/api.pharmacy.url", "")

	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return config, nil
}
