package app

import (
	"log"
	"os"

	"github.com/hammadmajid/gotic/internal/project"
)

type App struct {
	Logger  *log.Logger
	Project *project.Project
}

func New() (*App, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	project, err := project.New()
	if err != nil {
		return nil, err
	}

	return &App{
		Logger:  logger,
		Project: project,
	}, nil
}
