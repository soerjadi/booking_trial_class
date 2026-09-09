package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/infrastructure/config"
)

// pool backs the package-level Do/QuerierFromContext helpers in tx.go, so
// services and repositories can use them without any injected dependency.
var pool *pgxpool.Pool

// NewPGXPool creates a new PostgreSQL connection pool using the provided configuration
func NewPGXPool(ctx context.Context, cfg config.DBConfig) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(cfg.URL)
	if err != nil {
		log.ErrorCtx(ctx,
			"failed to parse connection string",
			log.Field("connection string", cfg.URL),
			log.Field("error", err),
		)
		return nil, err
	}

	// Apply pool settings from config
	poolConfig.MaxConns = int32(cfg.MaxConns)
	poolConfig.MinConns = int32(cfg.MinConns)
	poolConfig.MaxConnLifetime = time.Duration(cfg.MaxConnLifetime) * time.Minute
	poolConfig.MaxConnIdleTime = time.Duration(cfg.MaxConnIdleTime) * time.Minute

	pgPool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.ErrorCtx(ctx,
			"failed to create pool",
			log.Field("error", err),
		)
		return nil, err
	}

	// Ping the database to ensure connection is valid
	if err := pgPool.Ping(ctx); err != nil {
		log.ErrorCtx(ctx,
			"failed to ping database",
			log.Field("error", err),
		)
		return nil, err
	}

	log.InfoCtx(ctx, "database successfully connected")

	pool = pgPool
	return pgPool, nil
}
