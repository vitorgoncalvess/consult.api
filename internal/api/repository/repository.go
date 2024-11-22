package repository

import (
	"consult/internal/api/database"

	"github.com/spf13/viper"
)

type Repository struct {
	database *database.Database
}

func New(database *database.Database, config *viper.Viper) *Repository {
	return &Repository{database}
}
