package booking

import (
	"context"
	"testing"

	mock_repo "github.com/soerjadi/booking/internal/core/ports/repository/mocks"
	"github.com/soerjadi/booking/internal/infrastructure/db"
	"go.uber.org/mock/gomock"
)

func setupMocks(t *testing.T) (*gomock.Controller, *mock_repo.MockBookingRepositoryInterface, *mock_repo.MockTrialClassRepositoryInterface, *mock_repo.MockStudentRepositoryInterface, *mock_repo.MockPaymentAttemptRepositoryInterface) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_repo.NewMockBookingRepositoryInterface(ctrl)
	mockClassRepo := mock_repo.NewMockTrialClassRepositoryInterface(ctrl)
	mockStudentRepo := mock_repo.NewMockStudentRepositoryInterface(ctrl)
	mockPaymentRepo := mock_repo.NewMockPaymentAttemptRepositoryInterface(ctrl)
	return ctrl, mockRepo, mockClassRepo, mockStudentRepo, mockPaymentRepo
}

func init() {
	db.Do = func(ctx context.Context, fn func(ctx context.Context) error) error {
		return fn(ctx)
	}
}
