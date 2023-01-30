package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Server struct {
	Port string `mapstructure:"PORT"`
	Test string `mapstructure:"TEST"`
}

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
	Server Server `mapstructure:"SERVER"`
	Proxy Proxy `mapstructure:"PROXY"`
	GRPC GRPC `mapstructure:"GRPC"`
	Redis Redis `mapstructure:"REDIS"`
}

func LoadConfig() *Config {
	var config Config
	replacer := strings.NewReplacer(".", "_")
    	viper.SetEnvKeyReplacer(replacer)
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return &config
}
