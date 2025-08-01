package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewPostgres – подключение к PostgresSQL.
func NewPostgres(cfg Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
	)

	pool, err := pgxpool.New(context.Background(), dsn)

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Проверка подключения
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}
