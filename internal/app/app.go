package app

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/oemahdev/logger"
	"github.com/rs/cors"
	restmiddleware "github.com/soerjadi/booking/internal/adapters/primary/rest/middleware"
	restparent "github.com/soerjadi/booking/internal/adapters/primary/rest/parent"
	restrouter "github.com/soerjadi/booking/internal/adapters/primary/rest/router"
	parentsvc "github.com/soerjadi/booking/internal/core/services/parent"
	"github.com/soerjadi/booking/internal/infrastructure/config"
	"github.com/soerjadi/booking/internal/infrastructure/db"

	parentRepo "github.com/soerjadi/booking/internal/adapters/secondary/repository/parent"
)

type App struct {
	handler http.Handler
	port    int64
	dbPool  *pgxpool.Pool
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	pool, err := db.NewPGXPool(ctx, cfg.Database)
	if err != nil {
		log.ErrorCtx(ctx, "failed to create pool", log.Field("error", err))
		return nil, err
	}

	parentRepo := parentRepo.NewParentRepository(pool)

	parentSvc := parentsvc.NewParentService(parentRepo)
	parentHandler := restparent.NewHandler(parentSvc)

	r := mux.NewRouter()
	r.Use(restmiddleware.Logger)
	restrouter.RegisterRoutes(r, restrouter.Config{
		ParentHandler: parentHandler,
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   cfg.App.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})

	return &App{
		handler: c.Handler(r),
		port:    cfg.App.Port,
		dbPool:  pool,
	}, nil
}

func (a *App) Router() http.Handler { return a.handler }
func (a *App) Port() int64          { return a.port }

func (a *App) Shutdown(ctx context.Context) error {

	if a.dbPool != nil {
		a.dbPool.Close()
	}
	return nil
}
