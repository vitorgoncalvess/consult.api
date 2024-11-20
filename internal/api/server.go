package api

import (
	"consult/internal/api/handler"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Server struct {
	app    *echo.Echo
	config *viper.Viper
}

func New(config *viper.Viper) *Server {
	server := &Server{
		app:    echo.New(),
		config: config,
	}

	jwtSecret := []byte(config.GetString("jwt_secret"))

	server.routes(handler.New())

	return server
}

func (s *Server) Start() {
	if err := s.app.Start(":" + s.config.GetString("server.port")); err != nil {
		log.Fatalf("Erro ao inicializar a aplicação: %v", err)
	}
}
