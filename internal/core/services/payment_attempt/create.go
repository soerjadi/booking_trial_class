package paymentattempt

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

func (s *paymentAttemptService) Create(ctx context.Context, request domain.PaymentAttempt) (domain.PaymentAttempt, error) {
	return s.repo.Create(ctx, request)
}
