package booking

import "github.com/soerjadi/booking/internal/core/ports/repository"

type bookingRepository struct {
	db repository.DB
}

func NewBookingRepository(db repository.DB) repository.BookingRepositoryInterface {
	return &bookingRepository{db: db}
}
