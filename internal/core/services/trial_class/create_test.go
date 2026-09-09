package trialclass

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/soerjadi/booking/internal/core/domain"
	mock_repo "github.com/soerjadi/booking/internal/core/ports/repository/mocks"
)

func TestTrialClassService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockTrialClassRepositoryInterface(ctrl)
	service := NewTrialClassService(mockRepo)

	ctx := context.Background()
	req := domain.TrialClass{Name: "Math 101"}
	resp := domain.TrialClass{ID: 1, Name: "Math 101"}

	mockRepo.EXPECT().Create(ctx, req).Return(resp, nil)

	res, err := service.Create(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, resp, res)
}
