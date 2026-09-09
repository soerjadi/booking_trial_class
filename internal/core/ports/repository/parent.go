package repository

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

//go:generate mockgen -source=parent.go -destination=mocks/mock_parent.go -package=mocks
type ParentRepositoryInterface interface {
	Create(ctx context.Context, request domain.Parent) (domain.Parent, error)
	GetByID(ctx context.Context, request int64) (domain.Parent, error)
}
