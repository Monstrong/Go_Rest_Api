package main

import (
	"log/slog"
	"os"

	"github.com/Monstrong/Go_Rest_Api/internal/config"
)

const (
	envLocal  = "local"
	envDev    = "dev"
	envProd   = "prod"
)

func main() {
	// init config: cleanenv
	cfg := config.MustLoad()

	// TODO: init logger: slog
	log := setupLogger(cfg.Env)

	// log = log.With(slog.String("env", cfg.Env)) - тогда бы в каждое сообщение добавлялась инфа об окружении (у нас только в первой)
	log.Info("Starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// TODO: init storage: sqlite
	// TODO: init router: chi, render
	// TODO: run server
}

// setupLogger creates a logger depending on the given environment.
// It uses:
// - slog.TextHandler with LevelDebug for local development
// - slog.JSONHandler with LevelDebug for dev and prod environments
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}