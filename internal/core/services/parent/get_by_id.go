package parent

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

func (s *parentService) GetByID(ctx context.Context, id int64) (domain.Parent, error) {
	return s.repo.GetByID(ctx, id)
}
