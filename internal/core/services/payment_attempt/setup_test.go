package paymentattempt

import (
	"context"
	"testing"

	mock_repo "github.com/soerjadi/booking/internal/core/ports/repository/mocks"
	"github.com/soerjadi/booking/internal/infrastructure/db"
	"go.uber.org/mock/gomock"
)

func setupMocks(t *testing.T) (*gomock.Controller, *mock_repo.MockPaymentAttemptRepositoryInterface, *mock_repo.MockBookingRepositoryInterface, *mock_repo.MockTrialClassRepositoryInterface) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_repo.NewMockPaymentAttemptRepositoryInterface(ctrl)
	mockBookingRepo := mock_repo.NewMockBookingRepositoryInterface(ctrl)
	mockClassRepo := mock_repo.NewMockTrialClassRepositoryInterface(ctrl)
	return ctrl, mockRepo, mockBookingRepo, mockClassRepo
}

func init() {
	db.Do = func(ctx context.Context, fn func(ctx context.Context) error) error {
		return fn(ctx)
	}
}
