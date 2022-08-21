package config

import (
	"log"

	"github.com/spf13/viper"
)

func New(env string) *viper.Viper {
	config := viper.New()
	config.SetConfigType("env")
	config.SetConfigName(env)
	config.AddConfigPath(".")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Fatal("error on parsing configuration file")
		return nil
	}
	return config
}
