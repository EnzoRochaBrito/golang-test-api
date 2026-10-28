package main

import (
	repo "api/internal/adapters/postgresql/sqlc"
	"api/internal/json"
	"api/internal/note"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func (a *Application) Mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.Write(w, 200, "pong")	
	})

	noteService := note.NewService(repo.New(a.db))
	noteHandler := note.NewHandler(noteService)

	note.RegisterRoutes(mux, *noteHandler)
	return mux
}

func (a *Application) Run(m *http.ServeMux) {
	http.ListenAndServe(a.config.addr, m)
}

type Application struct {
	config Config
	db     *pgxpool.Pool
}

type Config struct {
	addr string
	db   DbConfig
}

type DbConfig struct {
	dsn string
}
