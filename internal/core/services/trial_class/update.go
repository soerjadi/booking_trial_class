package trialclass

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

func (s *trialClassService) Update(ctx context.Context, request domain.TrialClass) error {
	return s.repo.Update(ctx, request)
}
