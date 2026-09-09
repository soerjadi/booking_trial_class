package booking

import (
	"github.com/soerjadi/booking/internal/core/ports/repository"
	"github.com/soerjadi/booking/internal/core/ports/services"
)

type bookingService struct {
	repo        repository.BookingRepositoryInterface
	repoClass   repository.TrialClassRepositoryInterface
	repoStudent repository.StudentRepositoryInterface
	repoPayment repository.PaymentAttemptRepositoryInterface
}

func NewBookingService(
	repo repository.BookingRepositoryInterface,
	repoClass repository.TrialClassRepositoryInterface,
	repoStudent repository.StudentRepositoryInterface,
	repoPayment repository.PaymentAttemptRepositoryInterface,
) services.BookingServiceInterface {
	return &bookingService{
		repo:        repo,
		repoClass:   repoClass,
		repoStudent: repoStudent,
		repoPayment: repoPayment,
	}
}
