package main

import (
	"consult/internal/api"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(".")

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	config.SetEnvPrefix("APP")
	config.AutomaticEnv()

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Erro ao ler arquivo de configuração: %v", err)
	}

	server := api.New(config)

	server.Start()
}
