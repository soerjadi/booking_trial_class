package router

import (
	"github.com/gorilla/mux"
	restparent "github.com/soerjadi/booking/internal/adapters/primary/rest/parent"
)

type Config struct {
	ParentHandler *restparent.Handler
}

func RegisterRoutes(r *mux.Router, cfg Config) {
	v1 := r.PathPrefix("/v1").Subrouter()
	cfg.ParentHandler.Mount(v1)
}
