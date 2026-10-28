package main

import (
	"api/internal/env"
	"context"
	"log/slog"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	ctx := context.Background()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	err := godotenv.Load()
	if err != nil {
		logger.Warn("Warning: .env file not found")
	}

	config := Config{
		addr: ":8080",
		db: DbConfig{
			dsn: env.GetEnv("DATABASE_URL", ""),
		},
	}

	conn, err := pgxpool.New(ctx, config.db.dsn)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	logger.Info("database connected", "dsn", config.db.dsn)

	app := Application{
		config: config,
		db:     conn,
	}

	app.Run(app.Mount())

}
