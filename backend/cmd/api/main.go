package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"nexus/backend/internal/config"
	"nexus/backend/internal/repository"
	"nexus/backend/internal/server"
)

func main() {
	cfg := config.FromEnv()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	store, err := repository.NewStore(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("connect store: %v", err)
	}
	defer store.Close()

	if err := server.Run(ctx, cfg, store); err != nil {
		log.Fatalf("run server: %v", err)
	}
}
