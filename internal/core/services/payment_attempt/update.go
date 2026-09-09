package paymentattempt

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

func (s *paymentAttemptService) Update(ctx context.Context, request domain.PaymentAttempt) error {
	return s.repo.Update(ctx, request)
}
