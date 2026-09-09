package app

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/oemahdev/logger"
	"github.com/rs/cors"
	restbooking "github.com/soerjadi/booking/internal/adapters/primary/rest/booking"
	restmiddleware "github.com/soerjadi/booking/internal/adapters/primary/rest/middleware"
	restparent "github.com/soerjadi/booking/internal/adapters/primary/rest/parent"
	restpayment "github.com/soerjadi/booking/internal/adapters/primary/rest/payment"
	restrouter "github.com/soerjadi/booking/internal/adapters/primary/rest/router"
	reststudent "github.com/soerjadi/booking/internal/adapters/primary/rest/student"
	resttrialclass "github.com/soerjadi/booking/internal/adapters/primary/rest/trial_class"
	bookingsvc "github.com/soerjadi/booking/internal/core/services/booking"
	parentsvc "github.com/soerjadi/booking/internal/core/services/parent"
	paymentattemptsvc "github.com/soerjadi/booking/internal/core/services/payment_attempt"
	studentsvc "github.com/soerjadi/booking/internal/core/services/student"
	trialclasssvc "github.com/soerjadi/booking/internal/core/services/trial_class"
	"github.com/soerjadi/booking/internal/infrastructure/config"
	"github.com/soerjadi/booking/internal/infrastructure/db"

	bookingRepo "github.com/soerjadi/booking/internal/adapters/secondary/repository/booking"
	parentRepo "github.com/soerjadi/booking/internal/adapters/secondary/repository/parent"
	paymentAttemptRepo "github.com/soerjadi/booking/internal/adapters/secondary/repository/payment_attempt"
	studentRepo "github.com/soerjadi/booking/internal/adapters/secondary/repository/student"
	trialClassRepo "github.com/soerjadi/booking/internal/adapters/secondary/repository/trial_class"
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
	studentRepo := studentRepo.NewStudentRepository(pool)
	trialClassRepo := trialClassRepo.NewTrialClassRepository(pool)
	bookingRepo := bookingRepo.NewBookingRepository(pool)
	paymentAttemptRepo := paymentAttemptRepo.NewPaymentAttemptRepository(pool)

	parentSvc := parentsvc.NewParentService(parentRepo)
	studentSvc := studentsvc.NewStudentService(studentRepo)
	trialClassSvc := trialclasssvc.NewTrialClassService(trialClassRepo)
	bookingSvc := bookingsvc.NewBookingService(bookingRepo, trialClassRepo, studentRepo, paymentAttemptRepo)
	paymentAttemptSvc := paymentattemptsvc.NewPaymentAttemptService(paymentAttemptRepo, bookingRepo, trialClassRepo)

	parentHandler := restparent.NewHandler(parentSvc)
	studentHandler := reststudent.NewHandler(studentSvc)
	trialClassHandler := resttrialclass.NewHandler(trialClassSvc)
	bookingHandler := restbooking.NewHandler(bookingSvc)
	paymentHandler := restpayment.NewHandler(paymentAttemptSvc)

	r := mux.NewRouter()
	r.Use(restmiddleware.Logger)
	restrouter.RegisterRoutes(r, restrouter.Config{
		ParentHandler:     parentHandler,
		StudentHandler:    studentHandler,
		TrialClassHandler: trialClassHandler,
		BookingHandler:    bookingHandler,
		PaymentHandler:    paymentHandler,
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
