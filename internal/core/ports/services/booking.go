package services

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

//go:generate mockgen -source=booking.go -destination=mocks/mock_booking.go -package=mocks
type BookingServiceInterface interface {
	BookClass(ctx context.Context, request domain.BookCLassRequest) (domain.Booking, error)
	Update(ctx context.Context, request domain.Booking) error
}
