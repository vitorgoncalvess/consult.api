package main

import (
	"consult/internal/api"
	"log"

	"github.com/spf13/viper"
)

func main() {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(".")

	config.SetEnvPrefix("APP")
	config.BindEnv()

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Erro ao ler arquivo de configuração: %v", err)
	}

	server := api.New(config)

	server.Start()
}
