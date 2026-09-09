package paymentattempt

import (
	"github.com/soerjadi/booking/internal/core/ports/repository"
	"github.com/soerjadi/booking/internal/core/ports/services"
)

type paymentAttemptService struct {
	repo repository.PaymentAttemptRepositoryInterface
}

func NewPaymentAttemptService(repo repository.PaymentAttemptRepositoryInterface) services.PaymentAttemptServiceInterface {
	return &paymentAttemptService{repo: repo}
}
