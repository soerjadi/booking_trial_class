package booking

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

func (s *bookingService) Update(ctx context.Context, request domain.Booking) error {
	return s.repo.Update(ctx, request)
}
