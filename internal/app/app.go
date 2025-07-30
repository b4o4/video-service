package app

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"video-service/internal/config"
	"video-service/pkg/database"
)

type App struct {
	DB *pgxpool.Pool
}

func New() *App {
	// Загрузка конфига
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := createDB(cfg)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return &App{DB: db}
}

func (a *App) Run() {
	log.Println("Application started")
}

func createDB(cfg *config.Config) (*pgxpool.Pool, error) {
	// Инициализация Postgres
	db, err := database.NewPostgres(database.Config{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
	})

	return db, err
}
