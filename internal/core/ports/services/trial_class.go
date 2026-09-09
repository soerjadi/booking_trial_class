package services

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

//go:generate mockgen -source=trial_class.go -destination=mocks/mock_trial_class.go -package=mocks
type TrialClassServiceInterface interface {
	Create(ctx context.Context, request domain.TrialClass) (domain.TrialClass, error)
	Update(ctx context.Context, request domain.TrialClass) error
}
