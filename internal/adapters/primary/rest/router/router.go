package router

import (
	"github.com/gorilla/mux"
	restbooking "github.com/soerjadi/booking/internal/adapters/primary/rest/booking"
	restparent "github.com/soerjadi/booking/internal/adapters/primary/rest/parent"
	restpayment "github.com/soerjadi/booking/internal/adapters/primary/rest/payment"
	reststudent "github.com/soerjadi/booking/internal/adapters/primary/rest/student"
	resttrialclass "github.com/soerjadi/booking/internal/adapters/primary/rest/trial_class"
)

type Config struct {
	ParentHandler     *restparent.Handler
	StudentHandler    *reststudent.Handler
	TrialClassHandler *resttrialclass.Handler
	BookingHandler    *restbooking.Handler
	PaymentHandler    *restpayment.Handler
}

func RegisterRoutes(r *mux.Router, cfg Config) {
	v1 := r.PathPrefix("/v1").Subrouter()
	cfg.ParentHandler.Mount(v1)
	cfg.StudentHandler.Mount(v1)
	cfg.TrialClassHandler.Mount(v1)
	cfg.BookingHandler.Mount(v1)
	cfg.PaymentHandler.Mount(v1)
}
