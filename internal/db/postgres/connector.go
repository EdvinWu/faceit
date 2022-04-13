package postgres

import (
	"faceit-test/internal/config"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func Connect(cfg *config.Postgres) (*sqlx.DB, error) {
	c, err := sqlx.Connect("postgres", connectionString(cfg))
	if err != nil {
		return nil, err
	}
	return c, nil
}

func connectionString(cfg *config.Postgres) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database)
}
