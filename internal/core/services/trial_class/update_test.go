package trialclass

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/soerjadi/booking/internal/core/domain"
	mock_repo "github.com/soerjadi/booking/internal/core/ports/repository/mocks"
)

func TestTrialClassService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockTrialClassRepositoryInterface(ctrl)
	service := NewTrialClassService(mockRepo)

	ctx := context.Background()
	req := domain.TrialClass{ID: 1, Name: "Math 102"}

	mockRepo.EXPECT().Update(ctx, req).Return(nil)

	err := service.Update(ctx, req)
	assert.NoError(t, err)
}
