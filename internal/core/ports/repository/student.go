package repository

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

//go:generate mockgen -source=student.go -destination=mocks/mock_student.go -package=mocks
type StudentRepositoryInterface interface {
	Create(ctx context.Context, request domain.Student) (domain.Student, error)
	GetByID(ctx context.Context, request int64) (domain.Student, error)
}
