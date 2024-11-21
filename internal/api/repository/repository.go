package repository

import "consult/internal/api/database"

type Repository struct {
	database database.Database
}

func New(database database.Database) *Repository {
	return &Repository{database}
}
