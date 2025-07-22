package main

import (
	"log"

	"github.com/hammadmajid/gotic/internal/app"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatalf("Failed to initalize app")
	}

	app.Logger.Println("gotic")
}
