package booking

import (
	"github.com/soerjadi/booking/internal/core/ports/repository"
	"github.com/soerjadi/booking/internal/core/ports/services"
)

type bookingService struct {
	repo repository.BookingRepositoryInterface
}

func NewBookingService(repo repository.BookingRepositoryInterface) services.BookingServiceInterface {
	return &bookingService{repo: repo}
}
