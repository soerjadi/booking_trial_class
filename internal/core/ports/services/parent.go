package services

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

//go:generate mockgen -source=parent.go -destination=mocks/mock_parent.go -package=mocks
type ParentServiceInterface interface {
	Create(ctx context.Context, request domain.Parent) (domain.Parent, error)
	Update(ctx context.Context, request domain.Parent) error
}
