package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Conn interface {
	Query(query string, params ...any) (*sql.Rows, error)
	Exec(query string, params ...any) (sql.Result, error)
	QueryRow(query string, params ...any) *sql.Row
	Close() error
}

type Database struct {
	Conn Conn
}

func New(config *viper.Viper) *Database {
	return &Database{createConnection(config)}
}

func createConnection(config *viper.Viper) *sql.DB {
	strategy := config.GetString("server.database.strategy")

	switch strategy {
	case "mysql":
		return createMysqlConnection(config)
	default:
		return createMysqlConnection(config)
	}
}

func createMysqlConnection(config *viper.Viper) *sql.DB {
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

	return conn
}

func (d *Database) Close() {
	d.Conn.Close()
}
