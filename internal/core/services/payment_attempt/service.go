package paymentattempt

import (
	"github.com/soerjadi/booking/internal/core/ports/repository"
	"github.com/soerjadi/booking/internal/core/ports/services"
)

type paymentAttemptService struct {
	repo        repository.PaymentAttemptRepositoryInterface
	repoBooking repository.BookingRepositoryInterface
	repoClass   repository.TrialClassRepositoryInterface
}

func NewPaymentAttemptService(repo repository.PaymentAttemptRepositoryInterface, repoBooking repository.BookingRepositoryInterface, repoClass repository.TrialClassRepositoryInterface) services.PaymentAttemptServiceInterface {
	return &paymentAttemptService{
		repo:        repo,
		repoBooking: repoBooking,
		repoClass:   repoClass,
	}
}
