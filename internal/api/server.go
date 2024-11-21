package api

import (
	"consult/internal/api/database"
	"consult/internal/api/handler"
	middle "consult/internal/api/middleware"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

type Server struct {
	app      *echo.Echo
	config   *viper.Viper
	database *database.Database
}

func New(config *viper.Viper) *Server {
	server := &Server{
		app:      echo.New(),
		config:   config,
		database: database.New(config),
	}

	server.app.Use(middleware.Logger())
	server.app.Use(middleware.Recover())

	server.routes(handler.New(server.database, config), middle.New(config))

	return server
}

func (s *Server) Start() {
	if err := s.app.Start(":" + s.config.GetString("server.port")); err != nil {
		log.Fatalf("Erro ao inicializar a aplicação: %v", err)
	}
}

func (s *Server) Stop() {
	s.database.Close()
}
