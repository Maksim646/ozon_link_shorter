package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Maksim646/ozon_link_shorter/internal/app"
	"github.com/Maksim646/ozon_link_shorter/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"

	dataPostgres = "postgresql"
	dataMemory   = "inmemory"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.Any("config", cfg))

	// Create app instance
	application, cleanup, err := app.New(log, cfg)
	if err != nil {
		log.Error("failed to initialize app", slog.Any("error", err))
		os.Exit(1)
	}
	defer cleanup() // Ensure resources are released

	// Run gRPC server
	go application.GRPCSrv.MustRun()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	log.Info("stopping application", slog.String("signal", sign.String()))
	application.GRPCSrv.Stop()
	log.Info("application stopped")
}

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
