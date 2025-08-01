package migration

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"log"
	"video-service/pkg/database"
)

func Run(cfg database.Config) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
	)

	m, err := migrate.New(
		"file://migrations",
		dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
