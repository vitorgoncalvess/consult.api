package api

import (
	"consult/internal/api/handler"
	middle "consult/internal/api/middleware"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

type Server struct {
	app    *echo.Echo
	config *viper.Viper
	conn   *sql.DB
}

func New(config *viper.Viper) *Server {
	server := &Server{
		app:    echo.New(),
		config: config,
	}

	server.app.Use(middleware.Logger())
	server.app.Use(middleware.Recover())

	var (
		user     = config.GetString("MYSQL_USER")
		password = config.GetString("MYSQL_PASSWORD")
		host     = config.GetString("database.primary.host")
		port     = config.GetString("database.primary.port")
		name     = config.GetString("database.primary.name")
	)

	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name))

	if err != nil {
		log.Fatalf("Erro ao inicializar conexão com o banco: %v", err)
	}

	server.conn = conn

	server.routes(handler.New(conn, config), middle.New(config))

	return server
}

func (s *Server) Start() {
	if err := s.app.Start(":" + s.config.GetString("server.port")); err != nil {
		log.Fatalf("Erro ao inicializar a aplicação: %v", err)
	}
}

func (s *Server) Stop() {
	s.conn.Close()
}
