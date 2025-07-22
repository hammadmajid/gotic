package app

import (
	"log"
	"os"
)

type App struct {
	Logger *log.Logger
}

func New() (*App, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	return &App{
		Logger: logger,
	}, nil
}
