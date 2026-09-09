package app

import (
	"net/http"

	"github.com/gorilla/mux"
	restmiddleware "github.com/soerjadi/booking/internal/adapters/primary/rest/middleware"
	restrouter "github.com/soerjadi/booking/internal/adapters/primary/rest/router"
	restuser "github.com/soerjadi/booking/internal/adapters/primary/rest/user"
	usersvc "github.com/soerjadi/booking/internal/core/services/user"
)

type App struct {
	router *mux.Router
	port   string
}

func New() *App {
	svc := usersvc.NewDummyService()
	userHandler := restuser.NewHandler(svc)

	r := mux.NewRouter()
	r.Use(restmiddleware.Logger)
	restrouter.RegisterRoutes(r, restrouter.Config{
		UserHandler: userHandler,
	})

	return &App{
		router: r,
		port:   "8080",
	}
}

func (a *App) Router() http.Handler { return a.router }
func (a *App) Port() string         { return a.port }
