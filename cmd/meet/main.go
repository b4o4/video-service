package main

import "video-service/internal/app"

func main() {
	application := app.New()

	application.Run()
}
