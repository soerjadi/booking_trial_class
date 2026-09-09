package booking

import (
	"context"
	"fmt"

	"github.com/soerjadi/booking/internal/core/domain"
)

func (s *bookingService) BookClass(ctx context.Context, request domain.BookCLassRequest) (domain.Booking, error) {
	booking := domain.Booking{
		TrialClassID:   request.TrialClassID,
		StudentID:      request.StudentID,
		IdempotencyKey: fmt.Sprintf("%d", request.IdempotencyKey),
		Status:         domain.BookingStatusPending,
	}
	return s.repo.Create(ctx, booking)
}
