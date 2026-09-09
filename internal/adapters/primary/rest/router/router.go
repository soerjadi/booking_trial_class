package router

import (
	"github.com/gorilla/mux"
	restuser "github.com/soerjadi/booking/internal/adapters/primary/rest/user"
)

type Config struct {
	UserHandler *restuser.Handler
}

func RegisterRoutes(r *mux.Router, cfg Config) {
	v1 := r.PathPrefix("/v1").Subrouter()
	cfg.UserHandler.Mount(v1)
}
