package student

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

func (s *studentService) Create(ctx context.Context, request domain.Student) (domain.Student, error) {
	return s.repo.Create(ctx, request)
}
