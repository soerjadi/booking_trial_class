package booking

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/soerjadi/booking/internal/core/domain"
)

func TestBookingService_GetByStudentID(t *testing.T) {
	ctrl, mockRepo, mockClassRepo, mockStudentRepo, mockPaymentRepo := setupMocks(t)
	defer ctrl.Finish()

	service := NewBookingService(mockRepo, mockClassRepo, mockStudentRepo, mockPaymentRepo)

	ctx := context.Background()
	resp := []domain.Booking{{ID: 1, StudentID: 1}}

	mockRepo.EXPECT().GetByStudentID(ctx, int64(1)).Return(resp, nil)

	res, err := service.GetByStudentID(ctx, int64(1))
	assert.NoError(t, err)
	assert.Equal(t, resp, res)
}
