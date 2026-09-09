package parent

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

func (s *parentService) Create(ctx context.Context, request domain.Parent) (domain.Parent, error) {
	return s.repo.Create(ctx, request)
}
