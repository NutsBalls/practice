package migrations

import (
	"github.com/NutsBalls/practice-project-backend/internal/pkg/database"
	"github.com/pressly/goose"
)

func SetupMigrations(database database.DB) {
	goose.Up(database.Client, "internal/migrations")
}
