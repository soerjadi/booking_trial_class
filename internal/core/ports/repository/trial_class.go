package repository

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

//go:generate mockgen -source=trial_class.go -destination=mocks/mock_trial_class.go -package=mocks
type TrialClassRepositoryInterface interface {
	Create(ctx context.Context, request domain.TrialClass) (domain.TrialClass, error)
	Update(ctx context.Context, request domain.TrialClass) error
	InsertMember(ctx context.Context, request domain.TrialClassMember) error
}
