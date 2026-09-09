package trialclass

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

func (s *trialClassService) Create(ctx context.Context, request domain.TrialClass) (domain.TrialClass, error) {
	return s.repo.Create(ctx, request)
}
