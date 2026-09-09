package services

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

//go:generate mockgen -source=payment_attempt.go -destination=mocks/mock_payment_attempt.go -package=mocks
type PaymentAttemptServiceInterface interface {
	Create(ctx context.Context, request domain.PaymentAttempt) (domain.PaymentAttempt, error)
	Update(ctx context.Context, request domain.PaymentAttempt) error
}
