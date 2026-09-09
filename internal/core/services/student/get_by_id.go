package student

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

func (s *studentService) GetByID(ctx context.Context, id int64) (domain.Student, error) {
	return s.repo.GetByID(ctx, id)
}
