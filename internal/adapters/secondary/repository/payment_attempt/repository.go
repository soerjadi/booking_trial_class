package payment_attempt

import "github.com/soerjadi/booking/internal/core/ports/repository"

type paymentAttemptRepository struct {
	db repository.DB
}

func NewPaymentAttemptRepository(db repository.DB) repository.PaymentAttemptRepositoryInterface {
	return &paymentAttemptRepository{db: db}
}
