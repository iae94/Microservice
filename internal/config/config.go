package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port         uint `mapstructure:"Port"`
	Logger struct {
		Level            string   `mapstructure:"level"`
		Encoding         string   `mapstructure:"encoding"`
		OutputPaths      []string `mapstructure:"outputPaths"`
		ErrorOutputPaths []string `mapstructure:"errorOutputPaths"`
	} `mapstructure:"Logger"`
}

func ReadConfig() (conf *Config, err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("Reading config error: %v \n", err)
		return nil, err
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Printf("Unmarshaling config error: %v \n", err)
		return nil, err
	}
	return conf, nil
}
