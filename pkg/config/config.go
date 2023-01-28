package config

import (

	"github.com/spf13/viper"
)

type Redis struct{
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
	Password string `mapstructure:"PASSWORD"`
}

type Proxy struct {
	Chat string `mapstructure:"CHAT"`
	Conversation string `mapstructure:"CONVERSATION"`
}

type GRPC struct  {
	Account string `mapstructure:"ACCOUNT"`
}

type Config struct {
	Proxy Proxy `mapstructure:"PROXY"`
	GRPC GRPC `mapstructure:"GRPC"`
	Redis Redis `mapstructure:"REDIS"`
}

func LoadConfig() *Config {
	var config Config

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.ReadInConfig()

	err := viper.Unmarshal(&config)

	if err != nil {
		panic(err)
	}

	return &config
}
