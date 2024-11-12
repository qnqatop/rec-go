package app

import (
	"database/sql"
)

type App struct {
	db *sql.DB
}

func New(db *sql.DB) *App {
	return &App{db: db}
}

func (a *App) Run() error {
	a.registerHandlers()
	return a.runHTTPServer()
}
