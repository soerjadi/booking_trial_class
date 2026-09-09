package booking

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

func (s *bookingService) GetByStudentID(ctx context.Context, studentID int64) ([]domain.Booking, error) {
	return s.repo.GetByStudentID(ctx, studentID)
}
