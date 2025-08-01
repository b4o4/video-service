package main

import (
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"video-service/internal/app"
)

func main() {
	application := app.New()

	application.Run()
}
