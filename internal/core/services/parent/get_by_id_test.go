package parent

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/soerjadi/booking/internal/core/domain"
	mock_repo "github.com/soerjadi/booking/internal/core/ports/repository/mocks"
)

func TestParentService_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockParentRepositoryInterface(ctrl)
	service := NewParentService(mockRepo)

	ctx := context.Background()
	resp := domain.Parent{ID: 1, Name: "Jane Doe"}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetByID(ctx, int64(1)).Return(resp, nil)

		res, err := service.GetByID(ctx, int64(1))
		assert.NoError(t, err)
		assert.Equal(t, resp, res)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.EXPECT().GetByID(ctx, int64(2)).Return(domain.Parent{}, errors.New("not found"))

		_, err := service.GetByID(ctx, int64(2))
		assert.Error(t, err)
	})
}
