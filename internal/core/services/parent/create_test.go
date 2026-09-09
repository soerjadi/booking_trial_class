package parent

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/soerjadi/booking/internal/core/domain"
	mock_repo "github.com/soerjadi/booking/internal/core/ports/repository/mocks"
)

func TestParentService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockParentRepositoryInterface(ctrl)
	service := NewParentService(mockRepo)

	ctx := context.Background()
	req := domain.Parent{Name: "Jane Doe"}
	resp := domain.Parent{ID: 1, Name: "Jane Doe"}

	mockRepo.EXPECT().Create(ctx, req).Return(resp, nil)

	res, err := service.Create(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, resp, res)
}
