package services

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

type UserService interface {
	GetAll(ctx context.Context) ([]domain.User, error)
	GetByID(ctx context.Context, id int64) (*domain.User, error)
}
