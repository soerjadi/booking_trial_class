package services

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

//go:generate mockgen -source=booking.go -destination=mocks/mock_booking.go -package=mocks
type BookingServiceInterface interface {
	BookClass(ctx context.Context, request domain.BookClassRequest) (domain.Booking, error)
	Update(ctx context.Context, request domain.Booking) error
	GetByStudentID(ctx context.Context, studentID int64) ([]domain.Booking, error)
}
