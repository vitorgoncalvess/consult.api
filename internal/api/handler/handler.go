package handler

import (
	"database/sql"

	"github.com/spf13/viper"
)

type Conn interface {
	Query(query string, params ...any) (*sql.Rows, error)
	QueryRow(query string, params ...any) *sql.Row
}

type Handler struct {
	conn   Conn
	config *viper.Viper
}

func New(conn Conn, config *viper.Viper) *Handler {
	return &Handler{conn, config}
}
