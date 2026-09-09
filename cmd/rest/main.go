package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/app"
	"github.com/soerjadi/booking/internal/infrastructure/config"
)

func main() {
	cfg, err := config.Load("config.toml")
	if err != nil {
		log.Fatal("Failed to load config", log.Field("error", err))
	}

	log.Init(log.LogConfig{
		Level:      cfg.Log.Level,
		FilePath:   cfg.Log.FilePath,
		MaxSize:    cfg.Log.MaxSize,
		MaxBackups: cfg.Log.MaxBackups,
		MaxAge:     cfg.Log.MaxAge,
	})
	log.Info("Logger initialized successfully")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	a, err := app.New(ctx, cfg)
	if err != nil {
		log.Fatal("failed to initialize app", log.Field("error", err))
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.App.Port),
		Handler:      a.Router(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Info(fmt.Sprintf("Server listening on %s", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("Server error: %v", log.Field("error", err))
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	shutdownTimeout := time.Duration(cfg.App.ShutdownTimeout) * time.Second
	if shutdownTimeout == 0 {
		shutdownTimeout = 10 * time.Second
	}

	ctx, cancel = context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "HTTP server shutdown error: %v\n", err)
	}
	if err := a.Shutdown(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "App resources shutdown error: %v\n", err)
	}
	fmt.Println("Server stopped")
}
