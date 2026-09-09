package services

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

//go:generate mockgen -source=student.go -destination=mocks/mock_student.go -package=mocks
type StudentServiceInterface interface {
	Create(ctx context.Context, request domain.Student) (domain.Student, error)
	Update(ctx context.Context, request domain.Student) error
}
