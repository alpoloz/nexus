package server

import (
	"context"
	"net/http"
	"time"

	"nexus/backend/internal/api"
	"nexus/backend/internal/config"
	"nexus/backend/internal/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(ctx context.Context, cfg config.Config, store *repository.Store) error {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{cfg.FrontendURL},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodOptions},
	}))

	handler := api.NewHandler(store)
	e.GET("/healthz", handler.Health)
	e.GET("/api/blog-posts", handler.ListBlogPosts)
	e.GET("/api/startups", handler.ListStartups)

	errCh := make(chan error, 1)
	go func() {
		errCh <- e.Start(cfg.Address)
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return e.Shutdown(shutdownCtx)
	case err := <-errCh:
		return err
	}
}
